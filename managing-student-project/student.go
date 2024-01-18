package managingstudentproject

type Student struct {
	id         uint
	fullName   FullName
	group      string
	speciality string
}

func (s *Student) Id() uint {
	return s.id
}

func (s *Student) FullName() FullName {
	return s.fullName
}

func (s *Student) Group() string {
	return s.group
}

func (s *Student) Speciality() string {
	return s.speciality
}

func (s *Student) SetId(id uint) {
	s.id = id
}

func (s *Student) SetFullName(fullName FullName) {
	s.fullName = fullName
}

func (s *Student) SetGroup(group string) {
	s.group = group
}

func (s *Student) SetSpeciality(speciality string) {
	s.speciality = speciality
}

func InitStudent(id uint, name FullName, group string) Student {
	return Student{id: id, fullName: name, group: group}
}
