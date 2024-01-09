package managingstudentproject

type Student struct {
	fullName   FullName
	group      string
	speciality string
}

func InitStudent(name FullName, group string, spec string) Student {
	return Student{fullName: name, group: group, speciality: spec}
}
