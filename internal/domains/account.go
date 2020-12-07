package domains

import (
	"errors"
	"fmt"
	"strconv"
)

// Account is the struct that represents primary port for /account
type Account struct {
	ID              int64   `json:"id"`
	UUID            string  `json:"uuid"`
	DocumentNumber  string  `json:"document_number"`
	AvailableCredit float64 `json:"available_credit_limit"`
	TotalCredit     float64 `json:"total_credit_limit"`
}

// Verify checks if Account's inputs are valid
func (a *Account) Verify() (err error) {
	if a.DocumentNumber == "" {
		return errors.New("Document number cannot be empty")
	}
	if _, err := strconv.Atoi(a.DocumentNumber); err != nil {
		return fmt.Errorf("%v is not a document number", a.DocumentNumber)
	}
	if a.TotalCredit <= 0 {
		return errors.New("TotalCredit cannot be zero or negative")
	}
	return
}
