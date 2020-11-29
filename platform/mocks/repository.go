package mocks

import "github.com/arxdsilva/knab/internal/domains"

type Repository struct{}

func NewRepository() domains.SecondaryPort {
	return &Repository{}
}

func (r *Repository) CreateAccount(a *domains.Account) (err error) {
	return
}
