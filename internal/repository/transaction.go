package repository

import "github.com/arxdsilva/knab/internal/domains"

// Transaction is the repository data representation
type Transaction struct{}

// NewTransaction  retorna um tipo que implementa a porta secundaria
func NewTransaction() domains.TransactionService {
	return &Transaction{}
}

// CreateTransaction is the repository handler for the transaction creation workflow
func (t *Transaction) CreateTransaction(dt *domains.Transaction) (err error) {
	return
}
