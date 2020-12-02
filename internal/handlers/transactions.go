package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/arxdsilva/knab/internal/domains"
	"github.com/kpango/glg"
)

// CreateTransaction handler
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
		errAPI := errors.New("account_id could not be found")
		http.Error(w, errAPI.Error(), http.StatusNotFound)
		return
	}
	// not finished
}
