package managingstudentproject

type Account struct {
	id          uint
	accountName string
	email       string
}

func (a *Account) Id() uint {
	return a.id
}

func (a *Account) SetId(id uint) {
	a.id = id
}

func (a *Account) AccountName() string {
	return a.accountName
}

func (a *Account) Email() string {
	return a.email
}

func (a *Account) SetAccountName(accountName string) {
	a.accountName = accountName
}

func (a *Account) SetEmail(email string) {
	a.email = email
}

func InitAccount(id uint, name string, email string) Account {
	return Account{id: id, accountName: name, email: email}
}
