package domains

import "context"

type port struct {
	repo SecondaryPort
	ctx  context.Context
}

// NewService receives a Secondary Port of domain
// and instantiates a Primary Port
func NewService(repo SecondaryPort) PrimaryPort {
	return &port{repo, context.Background()}
}

func (p *port) CreateAccount(a *Account) (err error) {
	return p.repo.CreateAccount(a)
}

func (p *port) AccountByID(a *Account) (err error) {
	return p.repo.AccountByID(a)
}

func (p *port) IsRegistered(doc string) (r bool, err error) {
	return p.repo.IsRegistered(doc)
}
