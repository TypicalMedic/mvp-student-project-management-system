package managingstudentproject

import "fmt"

type FullName struct {
	name       string
	surname    string
	middleName string
}

func (f *FullName) Name() string {
	return f.name
}

func (f *FullName) Surname() string {
	return f.surname
}

func (f *FullName) MiddleName() string {
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

func (f *FullName) ToString() string {
	return fmt.Sprint(f.surname, " ", f.name, " ", f.middleName)
}
