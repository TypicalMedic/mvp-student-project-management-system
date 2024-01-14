package managingstudentproject

type Student struct {
	id         int
	fullName   FullName
	group      string
	speciality string
}

func (s *Student) GetId() int {
	return s.id
}

func (s *Student) GetFullName() FullName {
	return s.fullName
}

func (s *Student) GetGroup() string {
	return s.group
}

func (s *Student) GetSpeciality() string {
	return s.speciality
}

func (s *Student) SetId(id int) {
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

func InitStudent(name FullName, group string, spec string) Student {
	return Student{fullName: name, group: group, speciality: spec}
}
