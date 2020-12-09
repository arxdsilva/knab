package domains

type APIService interface {
	AccountService
	TransactionService
}

type AccountService interface {
	CreateAccount(a *Account) (err error)
	AccountByID(a *Account) (err error)
	IsIDRegistered(doc string) (r bool, err error)
	UpdateAvaliableLimit(accID int64, amount float64) (err error)
}

type TransactionService interface {
	HasLimitToTransaction(t *Transaction) (b bool, err error)
	CreateTransaction(t *Transaction) (err error)
	TransactionsWithBalance(accountID int64) (ts []Transaction, err error)
}
