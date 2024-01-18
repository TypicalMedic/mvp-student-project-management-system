package project

import (
	entities "mvp-student-project-management-system/managing-student-project"
)

func SetTasks(tasks []entities.Task) func(*entities.Project) {
	return func(p *entities.Project) {
		p.SetTasks(tasks)
	}
}
