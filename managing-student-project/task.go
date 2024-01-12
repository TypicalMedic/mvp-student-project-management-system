package managingstudentproject

import (
	"time"

	"github.com/google/uuid"
)

type TaskStatus int

const (
	NotStarted TaskStatus = iota
	InProgress
	Finished
)

type ProjectStage int

const (
	Analysis ProjectStage = iota
	Design
	Development
	Testing
	Deployment
)

type Task struct {
	id           uuid.UUID
	deadline     time.Time
	description  string
	status       TaskStatus
	projectStage ProjectStage
	taskFolder   Folder
}

func InitTask(deadline time.Time, desc string, status TaskStatus, pStage ProjectStage) Task {
	return Task{id: uuid.New(), deadline: deadline, description: desc, status: status, projectStage: pStage}
}

func (t *Task) ConnectDriveFolder(folder Folder) {
	t.taskFolder = folder
}

func (t *Task) GetDriveFolder() Folder {
	return t.taskFolder
}
