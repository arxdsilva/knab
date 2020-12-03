package handlers

import "github.com/arxdsilva/knab/internal/domains"

type HTTPPrimaryAdapter struct {
	service domains.APIService
}

func NewHTTPPrimaryAdapter(s domains.APIService) *HTTPPrimaryAdapter {
	return &HTTPPrimaryAdapter{s}
}
