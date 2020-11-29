package middlewares

import (
	"net/http"

	"github.com/arxdsilva/knab/internal/domains"
	"github.com/arxdsilva/knab/internal/handlers"
	"github.com/arxdsilva/knab/internal/repository/account"
	"github.com/gorilla/mux"
)

func RouterRegister(r *mux.Router) {
	rep := account.NewAccount()
	s := domains.NewService(rep)
	adapter := handlers.NewHTTPPrimaryAdapter(s)
	r.HandleFunc("/", adapter.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/accounts", adapter.CreateAccount).Methods(http.MethodPost)
	r.HandleFunc("/accounts/{account_id:[0-9]+}", adapter.GetAccount).Methods(http.MethodGet)
	//	r.HandleFunc("/accounts/:account_uuid", nil).Methods(http.MethodGet)
	r.HandleFunc("/transactions", adapter.CreateTransaction).Methods(http.MethodPost)
}
