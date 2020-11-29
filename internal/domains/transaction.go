package domains

import "github.com/nuveo/dbtime"

type Transaction struct {
	ID              string      `json:"id"`
	UUID            string      `json:"uuid"`
	AccountID       string      `json:"account_id"`
	OperationTypeID string      `json:"operation_type_id"`
	Amount          float64     `json:"amount"`
	EventDate       dbtime.Time `json:"event_date"`
}
