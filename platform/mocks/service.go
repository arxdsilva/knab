package mocks

import (
	"context"

	"github.com/arxdsilva/knab/internal/domains"
)

type portMock struct {
	repo domains.SecondaryPort
	ctx  context.Context
}

func NewService(repo domains.SecondaryPort) domains.PrimaryPort {
	return &portMock{repo, context.Background()}
}

func (p *portMock) CreateAccount(a *domains.Account) (err error) {
	return p.repo.CreateAccount(a)
}
func (p *portMock) AccountByID(a *domains.Account) (err error) {
	return p.repo.AccountByID(a)
}
func (p *portMock) IsRegistered(doc string) (r bool, err error) {
	return p.repo.IsRegistered(doc)
}
