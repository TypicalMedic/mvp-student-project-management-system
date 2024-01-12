package managingstudentproject

import "github.com/google/uuid"

type Student struct {
	id             uuid.UUID
	fullName       FullName
	group          string
	speciality     string
	repoMngAccount Account
}

func InitStudent(name FullName, group string, spec string, repoAcc Account) Student {
	return Student{id: uuid.New(), fullName: name, group: group, speciality: spec, repoMngAccount: repoAcc}
}
