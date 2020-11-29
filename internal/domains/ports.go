package domains

type PrimaryPort interface {
	CreateAccount(a *Account) (err error)
}

type SecondaryPort interface {
	CreateAccount(a *Account) (err error)
}
