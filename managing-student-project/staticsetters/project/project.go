package project

import (
	entities "mvp-student-project-management-system/managing-student-project"
)

func SetTasks(tasks []entities.Task) func(*entities.Project) {
	return func(p *entities.Project) {
		p.SetTasks(tasks)
	}
}

func SetDriveFolder(folder entities.Folder) func(*entities.Project) {
	return func(p *entities.Project) {
		p.SetDriveFolder(folder)
	}
}

func SetRepository(repo entities.Repository) func(*entities.Project) {
	return func(p *entities.Project) {
		p.SetRepository(repo)
	}
}

func SetStudent(stud entities.Student) func(*entities.Project) {
	return func(p *entities.Project) {
		p.SetStudent(stud)
	}
}

func SetSupervisor(sup entities.Professor) func(*entities.Project) {
	return func(p *entities.Project) {
		p.SetSupervisor(sup)
	}
}
