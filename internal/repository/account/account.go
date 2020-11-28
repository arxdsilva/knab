package account

import (
	"context"

	"github.com/arxdsilva/knab/internal/domains"
)

type Repository struct {
	ctx context.Context
}

// NewPGRepository  retorna um tipo que implementa a porta secundaria
func NewRepository() domains.SecondaryPort {
	return &Repository{}
}
