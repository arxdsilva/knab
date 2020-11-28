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
	c := &domains.Account{}
	if err := json.NewDecoder(r.Body).Decode(c); err != nil {
		glg.Error("[CreateAccount]", "(Decode)", err.Error())
		errAPI := errors.New("Could not parse request body")
		http.Error(w, errAPI.Error(), http.StatusBadRequest)
		return
	}

}
