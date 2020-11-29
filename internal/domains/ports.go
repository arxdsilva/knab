package domains

type PrimaryPort interface {
	CreateAccount(a *Account) (err error)
	AccountByID(a *Account) (err error)
	IsRegistered(doc string) (r bool, err error)
}

type SecondaryPort interface {
	CreateAccount(a *Account) (err error)
	AccountByID(a *Account) (err error)
	IsRegistered(doc string) (r bool, err error)
}
