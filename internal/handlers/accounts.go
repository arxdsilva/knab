package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/arxdsilva/knab/internal/domains"
	"github.com/kpango/glg"
)

func (a *HTTPPrimaryAdapter) GetAccount(w http.ResponseWriter, r *http.Request) {

}

func (a *HTTPPrimaryAdapter) CreateAccount(w http.ResponseWriter, r *http.Request) {
	acc := &domains.Account{}
	if err := json.NewDecoder(r.Body).Decode(acc); err != nil {
		glg.Error("[CreateAccount]", "(Decode)", err.Error())
		errAPI := errors.New("Could not parse request body")
		http.Error(w, errAPI.Error(), http.StatusBadRequest)
		return
	}
	if err = acc.Verify(); err != nil {
		glg.Error("[CreateAccount]", "(Verify)", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = a.service.CreateAccount(acc); err != nil {
		glg.Error("[CreateAccount]", "(service.CreateAccount)", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	glg.Info("[CreateAccount] success ", acc.UUID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(acc)
}
