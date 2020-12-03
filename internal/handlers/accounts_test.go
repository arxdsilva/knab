package handlers

import (
	"encoding/json"
	"fmt"
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

func Test_CreateAccountOK(t *testing.T) {
	mrep := mocks.NewRepository()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/accounts", adapter.CreateAccount).
		Methods("POST")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	body := strings.NewReader(`{"document_number":"1234"}`)
	resp, err := http.Post(server.URL+"/accounts", "", body)
	assert.Nil(t, err)
	assert.Equal(t, 201, resp.StatusCode)
	acc := &domains.Account{}
	err = json.NewDecoder(resp.Body).Decode(acc)
	assert.Nil(t, err)
	assert.Equal(t, "1234", acc.DocumentNumber)
	assert.Equal(t, int64(0), acc.ID)
}

func Test_CreateAccount_AlreadyRegistered(t *testing.T) {
	mrep := mocks.NewRepositoryRegistered()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/accounts", adapter.CreateAccount).
		Methods("POST")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	body := strings.NewReader(`{"document_number":"1234"}`)
	resp, err := http.Post(server.URL+"/accounts", "", body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusConflict, resp.StatusCode)
	bodyByte, errR := ioutil.ReadAll(resp.Body)
	assert.Nil(t, errR)
	assert.Contains(t, string(bodyByte), "has already been registered")
}

func Test_CreateAccount_AlreadyRegisteredError(t *testing.T) {
	mrep := mocks.NewRepositoryRegisteredError()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/accounts", adapter.CreateAccount).
		Methods("POST")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	body := strings.NewReader(`{"document_number":"1234"}`)
	resp, err := http.Post(server.URL+"/accounts", "", body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusConflict, resp.StatusCode)
	bodyByte, errR := ioutil.ReadAll(resp.Body)
	assert.Nil(t, errR)
	assert.Contains(t, string(bodyByte), "already been registered")
}

func Test_CreateAccount_InvalidBody(t *testing.T) {
	mrep := mocks.NewRepositoryRegisteredError()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/accounts", adapter.CreateAccount).
		Methods("POST")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	body := strings.NewReader(`12345512312`)
	resp, err := http.Post(server.URL+"/accounts", "", body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotAcceptable, resp.StatusCode)
	bodyByte, errR := ioutil.ReadAll(resp.Body)
	assert.Nil(t, errR)
	assert.Contains(t, string(bodyByte), "parse request body")
}

func Test_CreateAccount_InvalidDocumentNumber(t *testing.T) {
	mrep := mocks.NewRepositoryRegisteredError()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/accounts", adapter.CreateAccount).
		Methods("POST")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	body := strings.NewReader(`{}`)
	resp, err := http.Post(server.URL+"/accounts", "", body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	bodyByte, errR := ioutil.ReadAll(resp.Body)
	assert.Nil(t, errR)
	assert.Contains(t, string(bodyByte), "Document number cannot be empty")
}

func Test_CreateAccount_InvalidDocumentNumber_NaN(t *testing.T) {
	mrep := mocks.NewRepositoryRegisteredError()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/accounts", adapter.CreateAccount).
		Methods("POST")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	body := strings.NewReader(`{"document_number":"abcd12341asd"}`)
	resp, err := http.Post(server.URL+"/accounts", "", body)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	bodyByte, errR := ioutil.ReadAll(resp.Body)
	assert.Nil(t, errR)
	assert.Contains(t, string(bodyByte), "not a document number")
}

func Test_GetAccountByID_OK(t *testing.T) {
	mrep := mocks.NewRepository()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/accounts/{account_id}", adapter.GetAccountByID).
		Methods("GET")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	fmt.Println(server.Config)
	resp, err := http.Get(server.URL + "/accounts/1")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	acc := &domains.Account{}
	err = json.NewDecoder(resp.Body).Decode(acc)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), acc.ID)
}

func Test_GetAccountByID_NotFound(t *testing.T) {
	mrep := mocks.NewRepositoryAccByIDErr()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/accounts/{account_id}", adapter.GetAccountByID).
		Methods("GET")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	fmt.Println(server.Config)
	resp, err := http.Get(server.URL + "/accounts/1")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func Test_GetAccountByID_NaNumber(t *testing.T) {
	mrep := mocks.NewRepositoryAccByIDErr()
	adapter := NewHTTPPrimaryAdapter(mrep)
	r := mux.NewRouter()
	n := negroni.Classic()
	r.HandleFunc("/accounts/{account_id}", adapter.GetAccountByID).
		Methods("GET")
	n.UseHandler(r)
	server := httptest.NewServer(n)
	defer server.Close()
	fmt.Println(server.Config)
	resp, err := http.Get(server.URL + "/accounts/abc1as31s")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotAcceptable, resp.StatusCode)
}
