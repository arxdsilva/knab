package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/arxdsilva/knab/internal/domains"
	"github.com/gorilla/mux"
	"github.com/kpango/glg"
)

// GetAccountByID handler
// lets API users search accounts by passing an ID
//
// Responses:
//
// 200 OK
//
// 404 Not found
//
// 406 NotAcceptable (ID is not a number)
//
func (a *HTTPPrimaryAdapter) GetAccountByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["account_id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		glg.Error("[GetAccountByID]", "(Atoi)", err.Error())
		errAPI := fmt.Errorf("id '%s' is not a number", idStr)
		http.Error(w, errAPI.Error(), http.StatusNotAcceptable)
		return
	}
	acc := &domains.Account{ID: int64(idInt)}
	if err := a.service.AccountByID(acc); err != nil {
		glg.Error("[GetAccountByID]", "(service.AccountByID)", err.Error())
		errAPI := fmt.Errorf("Could not find account using the id '%s'", idStr)
		http.Error(w, errAPI.Error(), http.StatusNotFound)
		return
	}
	glg.Info("[GetAccountByID] success ", acc.UUID)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(acc)
}

// CreateAccount handler
// lets API users create accounts by passing a document_number as body
//
// Example body: `{"document_number":"1234"}`
//
// Responses:
//
// 201 Created
//
// 400 Bad Request (invalid number)
//
// 406 NotAcceptable (body problems)
//
// 409 Conflict (document number already registered)
//
// 500 Internal Server Error (could not create account)
//
func (a *HTTPPrimaryAdapter) CreateAccount(w http.ResponseWriter, r *http.Request) {
	acc := &domains.Account{}
	if err := json.NewDecoder(r.Body).Decode(acc); err != nil {
		glg.Error("[CreateAccount]", "(Decode)", err.Error())
		errAPI := errors.New("Could not parse request body")
		http.Error(w, errAPI.Error(), http.StatusNotAcceptable)
		return
	}
	if err := acc.Verify(); err != nil {
		glg.Error("[CreateAccount]", "(Verify)", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	registered, err := a.service.IsIDRegistered(acc.DocumentNumber)
	if err != nil {
		glg.Error("[CreateAccount]", "(IsIDRegistered)", err.Error())
		apiErr := errors.New("Internal server error")
		http.Error(w, apiErr.Error(), http.StatusInternalServerError)
		return
	}
	if registered {
		apiErr := fmt.Errorf("Document Number '%s' has already been registered",
			acc.DocumentNumber)
		http.Error(w, apiErr.Error(), http.StatusConflict)
		return
	}
	if err := a.service.CreateAccount(acc); err != nil {
		glg.Error("[CreateAccount]", "(service.CreateAccount)", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	glg.Info("[CreateAccount] success ", acc.UUID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(acc)
}
