package managingstudentproject

import (
	"time"
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
	deadline     time.Time
	description  string
	status       TaskStatus
	projectStage ProjectStage
}

func InitTask(deadline time.Time, desc string, status TaskStatus, pStage ProjectStage) Task {
	return Task{deadline: deadline, description: desc, status: status, projectStage: pStage}
}
