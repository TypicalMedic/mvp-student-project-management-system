package interfaces

import (
	entities "mvp-student-project-management-system/managing-student-project"
)

type IProjectDataStorage interface {
	SaveProject(entities.Project)
	DeleteProject(entities.Project)
	GetProfessorProjects(entities.Professor) []entities.Project
}

type IMeetingDataStorage interface {
	SaveMeeting(entities.Meeting)
	DeleteMeeting(entities.Meeting)
	UpdateMeeting(entities.Meeting)
}

type ITaskDataStorage interface {
	SaveTask(entities.Task)
	DeleteTask(entities.Task)
	UpdateTask(entities.Task)
	GetProjectTasks(entities.Project) []entities.Task
}
