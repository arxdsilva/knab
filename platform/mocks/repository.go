package mocks

import "github.com/arxdsilva/knab/internal/domains"

type Repository struct{}

func NewRepository() domains.SecondaryPort                         { return &Repository{} }
func (r *Repository) CreateAccount(a *domains.Account) (err error) { return }
func (r *Repository) AccountByID(a *domains.Account) (err error)   { return }
func (r *Repository) IsRegistered(doc string) (rp bool, err error) { return }

type RepositoryRegistered struct{}

func NewRepositoryRegistered() domains.SecondaryPort                         { return &RepositoryRegistered{} }
func (r *RepositoryRegistered) CreateAccount(a *domains.Account) (err error) { return }
func (r *RepositoryRegistered) AccountByID(a *domains.Account) (err error)   { return }
func (r *RepositoryRegistered) IsRegistered(doc string) (rp bool, err error) { return true, nil }
