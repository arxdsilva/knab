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
