package managingstudentproject

type Account struct {
	accountName string
	email       string
}

func (a *Account) GetAccountName() string {
	return a.accountName
}

func (a *Account) GetEmail() string {
	return a.email
}

func (a *Account) SetAccountName(accountName string) {
	a.accountName = accountName
}

func (a *Account) SetEmail(email string) {
	a.email = email
}

func InitRepoMngAccount(name string, email string) Account {
	return Account{accountName: name, email: email}
}
