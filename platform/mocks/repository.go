package mocks

import (
	"errors"

	"github.com/arxdsilva/knab/internal/domains"
)

type Repository struct{}

func NewRepository() domains.APIService                            { return &Repository{} }
func (r *Repository) CreateAccount(a *domains.Account) (err error) { return }
func (r *Repository) AccountByID(a *domains.Account) (err error) {
	a.ID = 1
	return
}
func (r *Repository) IsIDRegistered(doc string) (rp bool, err error) { return }
func (r *Repository) CreateTransaction(t *domains.Transaction) (err error) {
	t.ID = 1
	return
}

type RepositoryRegistered struct{}

func NewRepositoryRegistered() domains.APIService                              { return &RepositoryRegistered{} }
func (r *RepositoryRegistered) CreateAccount(a *domains.Account) (err error)   { return }
func (r *RepositoryRegistered) AccountByID(a *domains.Account) (err error)     { return }
func (r *RepositoryRegistered) IsIDRegistered(doc string) (rp bool, err error) { return true, nil }
func (r *RepositoryRegistered) CreateTransaction(t *domains.Transaction) (err error) {
	return errors.New("error")
}

type RepositoryRegisteredError struct{}

func NewRepositoryRegisteredError() domains.APIService                            { return &RepositoryRegisteredError{} }
func (r *RepositoryRegisteredError) CreateAccount(a *domains.Account) (err error) { return }
func (r *RepositoryRegisteredError) AccountByID(a *domains.Account) (err error)   { return }
func (r *RepositoryRegisteredError) IsIDRegistered(doc string) (rp bool, err error) {
	return true, nil
}
func (r *RepositoryRegisteredError) CreateTransaction(t *domains.Transaction) (err error) { return }

type RepositoryAccByIDErr struct{}

func NewRepositoryAccByIDErr() domains.APIService                            { return &RepositoryAccByIDErr{} }
func (r *RepositoryAccByIDErr) CreateAccount(a *domains.Account) (err error) { return }
func (r *RepositoryAccByIDErr) AccountByID(a *domains.Account) (err error) {
	return errors.New("acc id not found")
}
func (r *RepositoryAccByIDErr) IsIDRegistered(doc string) (rp bool, err error) {
	return true, errors.New("some error")
}
func (r *RepositoryAccByIDErr) CreateTransaction(t *domains.Transaction) (err error) { return }
