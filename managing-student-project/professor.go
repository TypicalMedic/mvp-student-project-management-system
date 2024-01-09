package managingstudentproject

type Professor struct {
	fullName FullName
	position string
	meetings []Meeting
	projects []Project
}

// add lower interface
func (p *Professor) SuperviseProject(project Project) {
	p.projects = append(p.projects, project)
}

// add lower interface
func (p *Professor) ArrangeMeeting(meeting Meeting) {
	p.meetings = append(p.meetings, meeting)
}

func InitProfessor(name FullName, position string) Professor {
	return Professor{fullName: name, position: position, meetings: []Meeting{}, projects: []Project{}}
}
