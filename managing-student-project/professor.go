package managingstudentproject

type Professor struct {
	id             int
	fullName       FullName
	position       string
	meetings       []Meeting
	calendar       Calendar
	repoMngAccount Account
}

func (p *Professor) GetFullName() FullName {
	return p.fullName
}

func (p *Professor) GetPosition() string {
	return p.position
}

func (p *Professor) GetMeetings() []Meeting {
	return p.meetings
}

func (p *Professor) GetCalendar() Calendar {
	return p.calendar
}

func (p *Professor) GetRepoMngAccount() Account {
	return p.repoMngAccount
}

func (p *Professor) SetFullName(fullName FullName) {
	p.fullName = fullName
}

func (p *Professor) SetPosition(position string) {
	p.position = position
}

func (p *Professor) SetMeetings(meetings []Meeting) {
	p.meetings = meetings
}

func (p *Professor) SetCalendar(calendar Calendar) {
	p.calendar = calendar
}

func (p *Professor) SetRepoMngAccount(repoMngAccount Account) {
	p.repoMngAccount = repoMngAccount
}

func InitProfessor(name FullName, position string, meetings []Meeting, cal Calendar, repoAcc Account) Professor {
	return Professor{fullName: name, position: position, meetings: meetings, calendar: cal, repoMngAccount: repoAcc}
}

func (p *Professor) GetId() int {
	return p.id
}

func (p *Professor) SetId(id int) {
	p.id = id
}

// type Professor struct {
// 	id       int
// 	fullName FullName
// 	position string
// 	meetings []Meeting
// 	projects []Project
// }

// func InitProfessor(name FullName, position string, meetings []Meeting, projects []Project) Professor {
// 	return Professor{fullName: name, position: position, meetings: meetings, projects: projects}
// }

// func (p *Professor) GetMeetings() []Meeting {
// 	return p.meetings
// }

// func (p *Professor) GetProjects() []Project {
// 	return p.projects
// }

// func (p *Professor) GetId() int {
// 	return p.id
// }
