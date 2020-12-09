package domains

import (
	"errors"

	"github.com/nuveo/dbtime"
)

// Transaction is the struct that holds the way that
// the API will comunicate with external devices
type Transaction struct {
	ID              int64       `json:"id"`
	UUID            string      `json:"uuid"`
	AccountID       int64       `json:"account_id"`
	OperationTypeID int         `json:"operation_type_id"`
	Amount          float64     `json:"amount"`
	Balance         float64     `json:"balance"`
	EventDate       dbtime.Time `json:"event_date"`
}

// iota (
// 	_ = 0
// 	OpCredito
// )

// Verify method asserts that the given transaction is in compliance with
// what knab needs as logic
func (t *Transaction) Verify() (err error) {
	if t.AccountID <= 0 {
		return errors.New("account_id cannot be zero or negative")
	}
	if !isOperation(t.OperationTypeID) {
		return errors.New("operation_type_id is invalid")
	}
	if t.Amount == float64(0) {
		return errors.New("amount cannot be zero")
	}
	if (t.OperationTypeID < 4) && (t.Amount > 0) {
		t.Amount = t.Amount * -1
		t.Balance = t.Amount * -1
	}
	return
}
