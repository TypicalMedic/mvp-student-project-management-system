package interfaces

import (
	entities "mvp-student-project-management-system/managing-student-project"
)

type IProjectDataStorage interface {
	SaveProject(entities.Project)
	DeleteProject(entities.Project)
	GetProfessorProjects(entities.Professor) []entities.Project
}
