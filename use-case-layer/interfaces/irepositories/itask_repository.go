package interfaces

import (
	entities "mvp-student-project-management-system/managing-student-project"
)

type ITaskRepository interface {
	SaveTask(entities.Task)
	DeleteTask(entities.Task)
	UpdateTask(entities.Task)
	GetProjectTasks(entities.Project) []entities.Task
}
