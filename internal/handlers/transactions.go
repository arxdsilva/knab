package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/arxdsilva/knab/internal/domains"
	"github.com/kpango/glg"
)

var (
	ErrStatusInternalServer = errors.New("Internal Server Error")
)

// CreateTransaction handler
// lets API users create accounts by passing a document_number as body
//
// Example body: `{"account_id":1,"operation_type_id":1,"amount":10}`
//
// Responses:
//
// 201 Created
//
// 400 Bad Request (invalid account/operation/amount)
//
// 404 NotFound (account_id wasnt found in db)
//
// 406 NotAcceptable (could not parse body)
//
// 500 Internal Server Error (could not create account)
//
func (a *HTTPPrimaryAdapter) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	t := &domains.Transaction{}
	if err := json.NewDecoder(r.Body).Decode(t); err != nil {
		glg.Error("[CreateTransaction]", "(Decode)", err.Error())
		errAPI := errors.New("Could not parse request body")
		http.Error(w, errAPI.Error(), http.StatusNotAcceptable)
		return
	}
	if err := t.Verify(); err != nil {
		glg.Error("[CreateTransaction]", "(Verify)", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	acc := &domains.Account{ID: t.AccountID}
	if err := a.service.AccountByID(acc); err != nil {
		glg.Error("[CreateTransaction]", "(service.AccountByID)", err.Error())
		errAPI := fmt.Errorf("account_id '%v' could not be found", t.AccountID)
		http.Error(w, errAPI.Error(), http.StatusNotFound)
		return
	}
	t.AccountID = acc.ID
	if t.OperationTypeID < 4 {
		canTransact, err := a.service.HasLimitToTransaction(t)
		if err != nil {
			glg.Error("[CreateTransaction]", "(service.HasLimitToTransaction)", err.Error())
			http.Error(w, ErrStatusInternalServer.Error(), http.StatusInternalServerError)
			return
		}
		if !canTransact {
			errAPI := errors.New("Account doesnt have avaliable limit")
			glg.Error("[CreateTransaction]", "(service.HasLimitToTransaction)", errAPI.Error())
			http.Error(w, errAPI.Error(), http.StatusUnauthorized)
			return
		}
	}
	if err := a.service.CreateTransaction(t); err != nil {
		glg.Error("[CreateTransaction]", "(service.CreateTransaction)", err.Error())
		http.Error(w, ErrStatusInternalServer.Error(), http.StatusInternalServerError)
		return
	}
	acc.UpdateAvaliableLimit(t.Amount)
	if err := a.service.UpdateAvaliableLimit(acc.ID, acc.AvailableCredit); err != nil {
		glg.Error("[CreateTransaction]", "(service.UpdateAvaliableLimit)", err.Error())
		http.Error(w, ErrStatusInternalServer.Error(), http.StatusInternalServerError)
		return
	}
	// fila > pago positivo XYZ
	// caso op =4
	// atualizar transacoes c/ balance negativo
	// passos:
	// pegar transacoes c/ balance negativos []list
	// atualizar 1:1 balance
	glg.Info("[CreateTransaction] success ", t.UUID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(t)
}
