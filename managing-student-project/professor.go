package managingstudentproject

type Professor struct {
	id             uint
	fullName       FullName
	position       string
	meetings       []Meeting
	projects       []Project
	calendar       Calendar
	repoMngAccount Account
}

func (p *Professor) Projects() []Project {
	return p.projects
}

func (p *Professor) SetProjects(projects []Project) {
	p.projects = projects
}

func (p *Professor) FullName() FullName {
	return p.fullName
}

func (p *Professor) Position() string {
	return p.position
}

func (p *Professor) Meetings() []Meeting {
	return p.meetings
}

func (p *Professor) Calendar() Calendar {
	return p.calendar
}

func (p *Professor) RepoMngAccount() Account {
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

func InitProfessor(id uint, name FullName, position string, meetings []Meeting, projects []Project, cal Calendar, repoAcc Account) Professor {
	return Professor{id: id, fullName: name, position: position, meetings: meetings, projects: projects, calendar: cal, repoMngAccount: repoAcc}
}

func (p *Professor) Id() uint {
	return p.id
}

func (p *Professor) SetId(id uint) {
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

// func (p *Professor) Meetings() []Meeting {
// 	return p.meetings
// }

// func (p *Professor) Projects() []Project {
// 	return p.projects
// }

// func (p *Professor) Id() int {
// 	return p.id
// }
