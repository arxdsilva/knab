package repository

import (
	"context"

	"github.com/arxdsilva/knab/internal/domains"
	"github.com/arxdsilva/knab/platform/config"
	"github.com/nuveo/dbtime"
)

type Account struct {
	ctx            context.Context
	ID             int64       `json:"id"`
	UUID           string      `json:"uuid"`
	DocumentNumber string      `json:"document_number"`
	CreatedAt      dbtime.Time `json:"created_at"`
	Active         bool        `json:"active"`
}

// NewAccount  retorna um tipo que implementa a porta secundaria
func NewAccount() domains.SecondaryPort {
	return &Account{}
}

func (ar *Account) CreateAccount(a *domains.Account) (err error) {
	sql := `INSERT INTO accounts (document_number) VALUES ($1)`
	sc := config.Get.DBAdapter.Insert(sql, a.DocumentNumber)
	err = sc.Err()
	if err != nil {
		return
	}
	_, err = sc.Scan(ar)
	if err != nil {
		return
	}
	a.ID = ar.ID
	a.UUID = ar.UUID
	return
}

func (ar *Account) AccountByID(a *domains.Account) (err error) {
	sql := `SELECT id, uuid, document_number, created_at FROM accounts WHERE id=$1`
	sc := config.Get.DBAdapter.Query(sql, a.ID)
	err = sc.Err()
	if err != nil {
		return
	}
	_, err = sc.Scan(ar)
	if err != nil {
		return
	}
	a.ID = ar.ID
	a.UUID = ar.UUID
	a.DocumentNumber = ar.DocumentNumber
	return
}
