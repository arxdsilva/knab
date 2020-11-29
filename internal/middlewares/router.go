package middlewares

import (
	"net/http"

	"github.com/arxdsilva/knab/internal/domains"
	"github.com/arxdsilva/knab/internal/handlers"
	"github.com/arxdsilva/knab/internal/repository"
	"github.com/gorilla/mux"
)

func RouterRegister(r *mux.Router) {
	rep := repository.NewAccount()
	s := domains.NewService(rep)
	adapter := handlers.NewHTTPPrimaryAdapter(s)
	// to avoid prest dbname/table collision I had to create this path
	r.HandleFunc("/", adapter.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/prest/api/accounts", adapter.CreateAccount).Methods(http.MethodPost)
	r.HandleFunc("/prest/api/accounts/{account_id:[0-9]+}", adapter.GetAccountByID).Methods(http.MethodGet)
	//	r.HandleFunc("/accounts/:account_uuid", nil).Methods(http.MethodGet)
	r.HandleFunc("/prest/api/transactions", adapter.CreateTransaction).Methods(http.MethodPost)
}
