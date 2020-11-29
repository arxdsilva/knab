package repository

import (
	"context"

	"github.com/arxdsilva/knab/internal/domains"
	"github.com/arxdsilva/knab/platform/config"
)

type Account struct {
	ctx            context.Context
	ID             int64  `json:"id"`
	UUID           string `json:"uuid"`
	DocumentNumber string `json:"document_number"`
	Active         bool   `json:"active"`
}

// NewAccount  retorna um tipo que implementa a porta secundaria
func NewAccount() domains.SecondaryPort {
	return &Account{}
}

func (ar *Account) CreateAccount(a *domains.Account) (err error) {
	sql := `INSERT INTO accounts (document_number) VALUES ($1)`
	sc := config.Get.DBAdapter.Insert(sql, a.DocumentNumber)
	return
}
