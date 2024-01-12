package interfaces

import (
	entities "mvp-student-project-management-system/managing-student-project"
)

type IPresenter interface {
	ShowProjectsList([]entities.Project)
	ShowProject(entities.Project)
	ShowCommits([]entities.Commit)
	ShowSchedule([]entities.Meeting)
	ShowTasks([]entities.Task)
}
