package managingstudentproject

type FullName struct {
	name       string
	surname    string
	middleName string
}

func InitFullName(name string, surname string, middlename string) FullName {
	return FullName{name: name, surname: surname, middleName: middlename}
}
