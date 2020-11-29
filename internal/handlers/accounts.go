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

func (a *HTTPPrimaryAdapter) GetAccountByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["account_id"]
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		glg.Error("[GetAccountByID]", "(Atoi)", err.Error())
		errAPI := fmt.Errorf("id '%s' is not a number", idStr)
		http.Error(w, errAPI.Error(), http.StatusBadRequest)
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

func (a *HTTPPrimaryAdapter) CreateAccount(w http.ResponseWriter, r *http.Request) {
	acc := &domains.Account{}
	if err := json.NewDecoder(r.Body).Decode(acc); err != nil {
		glg.Error("[CreateAccount]", "(Decode)", err.Error())
		errAPI := errors.New("Could not parse request body")
		http.Error(w, errAPI.Error(), http.StatusBadRequest)
		return
	}
	if err := acc.Verify(); err != nil {
		glg.Error("[CreateAccount]", "(Verify)", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
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
