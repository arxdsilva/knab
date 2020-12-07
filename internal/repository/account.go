package repository

import (
	"fmt"

	"github.com/arxdsilva/knab/internal/domains"
	"github.com/arxdsilva/knab/platform/config"
	"github.com/nuveo/dbtime"
)

type Account struct {
	ID              int64       `json:"id"`
	UUID            string      `json:"uuid"`
	DocumentNumber  string      `json:"document_number"`
	CreatedAt       dbtime.Time `json:"created_at"`
	Active          bool        `json:"active"`
	AvailableCredit float64     `json:"available_credit_limit"`
	TotalCredit     float64     `json:"total_credit_limit"`
}

// NewAccount  retorna um tipo que implementa a porta secundaria
func NewAccount() domains.AccountService {
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
	a.AvailableCredit = ar.AvailableCredit
	a.TotalCredit = ar.TotalCredit
	return
}

func (ar *Account) AccountByID(a *domains.Account) (err error) {
	sql := `SELECT id, uuid, document_number, created_at FROM accounts WHERE id=$1`
	sc := config.Get.DBAdapter.Query(sql, a.ID)
	err = sc.Err()
	if err != nil {
		return
	}
	n, err := sc.Scan(ar)
	if err != nil {
		return
	}
	if n == 0 {
		return fmt.Errorf("Account id '%v' could not be found", a.ID)
	}
	a.ID = ar.ID
	a.UUID = ar.UUID
	a.DocumentNumber = ar.DocumentNumber
	return
}

func (ar *Account) IsIDRegistered(doc string) (r bool, err error) {
	sql := `SELECT id FROM accounts WHERE document_number=$1`
	sc := config.Get.DBAdapter.Query(sql, doc)
	err = sc.Err()
	if err != nil {
		return
	}
	n, errS := sc.Scan(ar)
	if errS != nil {
		return r, errS
	}
	if n > 0 {
		return true, nil
	}
	return
}
