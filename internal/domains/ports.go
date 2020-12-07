package domains

type APIService interface {
	AccountService
	TransactionService
}

type AccountService interface {
	CreateAccount(a *Account) (err error)
	AccountByID(a *Account) (err error)
	IsIDRegistered(doc string) (r bool, err error)
}

type TransactionService interface {
	HasLimitToTransaction(t *Transaction) (b bool, err error)
	CreateTransaction(t *Transaction) (err error)
}
