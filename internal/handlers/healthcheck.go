package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/arxdsilva/knab/internal/domains"
)

type HTTPPrimaryAdapter struct {
	service domains.APIService
}

func NewHTTPPrimaryAdapter(s domains.APIService) *HTTPPrimaryAdapter {
	return &HTTPPrimaryAdapter{s}
}

func (a *HTTPPrimaryAdapter) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ok := struct {
		Service string `json:"service"`
	}{"ok"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ok)
}
