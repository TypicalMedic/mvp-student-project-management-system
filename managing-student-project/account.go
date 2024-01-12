package managingstudentproject

type Account struct {
	accountName string
	email       string
}

func InitRepoMngAccount(name string, email string) Account {
	return Account{accountName: name, email: email}
}
