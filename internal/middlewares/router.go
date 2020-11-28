package middlewares

import (
	"net/http"

	"github.com/arxdsilva/knab/internal/handlers"
	"github.com/gorilla/mux"
)

func RouterRegister(r *mux.Router) {
	r.HandleFunc("/", handlers.HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/accounts", nil).Methods(http.MethodPost)
	r.HandleFunc("/accounts/{account_id:[0-9]+}", nil).Methods(http.MethodGet)
	//	r.HandleFunc("/accounts/:account_uuid", nil).Methods(http.MethodGet)
	r.HandleFunc("/transactions", nil).Methods("POST")
}
