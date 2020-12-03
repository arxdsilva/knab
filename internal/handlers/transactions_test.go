package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/arxdsilva/knab/internal/domains"
	"github.com/arxdsilva/knab/platform/mocks"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"
)

func Test_CreateTransaction_OK(t *testing.T) {
	mrep := mocks.NewRepository()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/transactions", adapter.CreateTransaction).
		Methods("POST")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	body := strings.NewReader(`{"account_id":1,"operation_type_id":1,"amount":10}`)
	resp, err := http.Post(server.URL+"/transactions", "", body)
	assert.Nil(t, err)
	assert.Equal(t, 201, resp.StatusCode)
	dt := &domains.Transaction{}
	err = json.NewDecoder(resp.Body).Decode(dt)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), dt.ID)
	assert.Equal(t, int64(1), dt.AccountID)
	assert.Equal(t, float64(-10), dt.Amount)
}

func Test_CreateTransaction_AccountNotFound(t *testing.T) {
	mrep := mocks.NewRepositoryAccByIDErr()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/transactions", adapter.CreateTransaction).
		Methods("POST")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	body := strings.NewReader(`{"account_id":1,"operation_type_id":1,"amount":10}`)
	resp, err := http.Post(server.URL+"/transactions", "", body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	bodyString := string(bodyBytes)
	assert.Contains(t, bodyString, "could not be found")
}

func Test_CreateTransaction_NotParseableBody(t *testing.T) {
	mrep := mocks.NewRepositoryAccByIDErr()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/transactions", adapter.CreateTransaction).
		Methods("POST")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	body := strings.NewReader(``)
	resp, err := http.Post(server.URL+"/transactions", "", body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotAcceptable, resp.StatusCode)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	bodyString := string(bodyBytes)
	assert.Contains(t, bodyString, "parse request body")
}

func Test_CreateTransaction_InvalidOperation(t *testing.T) {
	mrep := mocks.NewRepositoryAccByIDErr()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/transactions", adapter.CreateTransaction).
		Methods("POST")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	body := strings.NewReader(`{"account_id":1,"operation_type_id":10,"amount":10}`)
	resp, err := http.Post(server.URL+"/transactions", "", body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	bodyString := string(bodyBytes)
	assert.Contains(t, bodyString, "operation_type_id is invalid")
}

func Test_CreateTransaction_InvalidAccountID(t *testing.T) {
	mrep := mocks.NewRepositoryAccByIDErr()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/transactions", adapter.CreateTransaction).
		Methods("POST")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	body := strings.NewReader(`{"operation_type_id":1,"amount":10}`)
	resp, err := http.Post(server.URL+"/transactions", "", body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	bodyString := string(bodyBytes)
	assert.Contains(t, bodyString, "cannot be zero or nega")
}

func Test_CreateTransaction_InternalError(t *testing.T) {
	mrep := mocks.NewRepositoryRegistered()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/transactions", adapter.CreateTransaction).
		Methods("POST")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	body := strings.NewReader(`{"account_id":1,"operation_type_id":1,"amount":10}`)
	resp, err := http.Post(server.URL+"/transactions", "", body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	bodyString := string(bodyBytes)
	assert.Contains(t, bodyString, "Server Error")
}
