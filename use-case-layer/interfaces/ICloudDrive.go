package interfaces

import (
	entities "mvp-student-project-management-system/managing-student-project"
)

type ICloudDrive interface {
	CteateProjectFolder(entities.Project) entities.Folder
	CteateTaskFolder(task entities.Task, parentFolder entities.Folder) entities.Folder
	DeleteFolder(entities.Folder)
	CreateTaskFile(entities.Folder, entities.Task)
}
