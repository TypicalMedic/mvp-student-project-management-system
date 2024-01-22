package professor

import (
	entities "mvp-student-project-management-system/managing-student-project"
)

func SetMeetings(meetings []entities.Meeting) func(*entities.Professor) {
	return func(p *entities.Professor) {
		p.SetMeetings(meetings)
	}
}

func SetProjects(projcets []entities.Project) func(*entities.Professor) {
	return func(p *entities.Professor) {
		p.SetProjects(projcets)
	}
}

func SetCalendar(cal entities.Calendar) func(*entities.Professor) {
	return func(p *entities.Professor) {
		p.SetCalendar(cal)
	}
}

func SetRepoMngAccount(acc entities.Account) func(*entities.Professor) {
	return func(p *entities.Professor) {
		p.SetRepoMngAccount(acc)
	}
}

func SetDriveFolder(folder entities.Folder) func(*entities.Professor) {
	return func(p *entities.Professor) {
		p.SetDriveFolder(folder)
	}
}
