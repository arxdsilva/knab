package domains

type port struct {
	account     AccountService
	transaction TransactionService
}

// NewService receives a Secondary Port of domain
// and instantiates a Primary Port
func NewService(a AccountService, t TransactionService) APIService {
	return &port{account: a, transaction: t}
}

func (p *port) CreateAccount(a *Account) (err error) {
	return p.account.CreateAccount(a)
}

func (p *port) AccountByID(a *Account) (err error) {
	return p.account.AccountByID(a)
}

func (p *port) IsIDRegistered(doc string) (r bool, err error) {
	return p.account.IsIDRegistered(doc)
}

func (p *port) CreateTransaction(t *Transaction) (err error) {
	return p.transaction.CreateTransaction(t)
}

func (p *port) HasLimitToTransaction(t *Transaction) (b bool, err error) {
	return p.transaction.HasLimitToTransaction(t)
}

func (p *port) UpdateAvaliableLimit(accID int64, amount float64) (err error) {
	return p.account.UpdateAvaliableLimit(accID, amount)
}

func (p *port) TransactionsWithBalance(accID int64) (ts []Transaction, err error) {
	return p.transaction.TransactionsWithBalance(accID)
}
