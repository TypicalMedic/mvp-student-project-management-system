package managingstudentproject

import "github.com/google/uuid"

type Professor struct {
	id       uuid.UUID
	fullName FullName
	position string
	meetings []Meeting
	calendar Calendar
}

func InitProfessor(name FullName, position string, meetings []Meeting, cal Calendar) Professor {
	return Professor{id: uuid.New(), fullName: name, position: position, meetings: meetings, calendar: cal}
}

func (p *Professor) GetMeetings() []Meeting {
	return p.meetings
}

func (p *Professor) GetId() uuid.UUID {
	return p.id
}

// type Professor struct {
// 	id       uuid.UUID
// 	fullName FullName
// 	position string
// 	meetings []Meeting
// 	projects []Project
// }

// func InitProfessor(name FullName, position string, meetings []Meeting, projects []Project) Professor {
// 	return Professor{id: uuid.New(), fullName: name, position: position, meetings: meetings, projects: projects}
// }

// func (p *Professor) GetMeetings() []Meeting {
// 	return p.meetings
// }

// func (p *Professor) GetProjects() []Project {
// 	return p.projects
// }

// func (p *Professor) GetId() uuid.UUID {
// 	return p.id
// }
