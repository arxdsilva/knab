package domains

import (
	"errors"

	"github.com/nuveo/dbtime"
)

type Transaction struct {
	ID              string      `json:"id"`
	UUID            string      `json:"uuid"`
	AccountID       string      `json:"account_id"`
	OperationTypeID int         `json:"operation_type_id"`
	Amount          float64     `json:"amount"`
	EventDate       dbtime.Time `json:"event_date"`
}

func (t *Transaction) Verify() (err error) {
	if t.AccountID == "" {
		return errors.New("account_id cannot be empty")
	}
	if !IsOperation(t.OperationTypeID) {
		return errors.New("operation_type_id is invalid")
	}
	if t.Amount == float64(0) {
		return errors.New("amount cannot be zero")
	}
	if (t.OperationTypeID < 4) && (t.Amount > 0) {
		t.Amount = t.Amount * -1
	}
	return
}
