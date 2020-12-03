package handlers

import (
	"encoding/json"
	"net/http"
)

func (a *HTTPPrimaryAdapter) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ok := struct {
		Service string `json:"service"`
	}{"ok"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ok)
}
