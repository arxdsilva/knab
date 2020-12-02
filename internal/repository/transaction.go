package repository

import (
	"github.com/arxdsilva/knab/internal/domains"
	"github.com/arxdsilva/knab/platform/config"
	"github.com/nuveo/dbtime"
)

// Transaction is the repository data representation
type Transaction struct {
	ID          int64       `json:"id"`
	AccountID   int64       `json:"account_id"`
	OperationID int64       `json:"operation_id"`
	Amount      float64     `json:"amount"`
	UUID        string      `json:"uuid"`
	EventDate   dbtime.Time `json:"event_date"`
	LastUpdate  dbtime.Time `json:"last_update"`
}

// NewTransaction  retorna um tipo que implementa a porta secundaria
func NewTransaction() domains.TransactionService {
	return &Transaction{}
}

// CreateTransaction is the repository handler for the transaction creation workflow
func (t *Transaction) CreateTransaction(dt *domains.Transaction) (err error) {
	sql := `INSERT INTO accounts (account_id, operation_id, amount) 
		VALUES ($1, $2, $3)`
	sc := config.Get.DBAdapter.Insert(sql,
		dt.AccountID, dt.OperationTypeID, dt.Amount)
	err = sc.Err()
	if err != nil {
		return
	}
	_, err = sc.Scan(t)
	if err != nil {
		return
	}
	dt.ID = t.ID
	dt.UUID = t.UUID
	dt.EventDate = t.EventDate
	return
}
