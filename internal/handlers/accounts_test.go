package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/arxdsilva/knab/internal/domains"
	"github.com/arxdsilva/knab/platform/config"
	"github.com/arxdsilva/knab/platform/mocks"
	"github.com/gorilla/mux"
	"github.com/prest/prest/adapters/mock"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/negroni"
)

func init() {
	testing.Init()
	config.Load()
}

func Test_CreateAccountOK(t *testing.T) {
	os.Setenv("TEST", "true")
	defer os.Unsetenv("TEST")
	m := mock.New(t)
	config.Get.DBAdapter = m
	mrep := mocks.NewRepository()
	// s := mocks.NewService(mrep)
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
	os.Setenv("TEST", "true")
	defer os.Unsetenv("TEST")
	m := mock.New(t)
	config.Get.DBAdapter = m
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
	os.Setenv("TEST", "true")
	defer os.Unsetenv("TEST")
	m := mock.New(t)
	config.Get.DBAdapter = m
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
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	bodyByte, errR := ioutil.ReadAll(resp.Body)
	assert.Nil(t, errR)
	assert.Contains(t, string(bodyByte), "server error")
}

func Test_CreateAccount_InvalidBody(t *testing.T) {
	os.Setenv("TEST", "true")
	defer os.Unsetenv("TEST")
	m := mock.New(t)
	config.Get.DBAdapter = m
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

func Test_GetAccountByID_OK(t *testing.T) {
	os.Setenv("TEST", "true")
	defer os.Unsetenv("TEST")
	m := mock.New(t)
	config.Get.DBAdapter = m
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
	resp, err := http.Get(server.URL + "/accounts/" + "1")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	acc := &domains.Account{}
	err = json.NewDecoder(resp.Body).Decode(acc)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), acc.ID)
}
