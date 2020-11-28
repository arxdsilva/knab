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
