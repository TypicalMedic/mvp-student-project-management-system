package managingstudentproject

type FullName struct {
	name       string
	surname    string
	middleName string
}

func (f *FullName) GetName() string {
	return f.name
}

func (f *FullName) GetSurname() string {
	return f.surname
}

func (f *FullName) GetMiddleName() string {
	return f.middleName
}

func (f *FullName) SetName(name string) {
	f.name = name
}

func (f *FullName) SetSurname(surname string) {
	f.surname = surname
}

func (f *FullName) SetMiddleName(middleName string) {
	f.middleName = middleName
}

func InitFullName(name string, surname string, middlename string) FullName {
	return FullName{name: name, surname: surname, middleName: middlename}
}
