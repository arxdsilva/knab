package repository

import "github.com/arxdsilva/knab/internal/domains"

type Transaction struct{}

// NewTransaction  retorna um tipo que implementa a porta secundaria
func NewTransaction() domains.TransactionService {
	return &Transaction{}
}

func (t *Transaction) CreateTransaction(dt *domains.Transaction) (err error) {
	return
}
