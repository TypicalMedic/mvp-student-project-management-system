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
	projectSrage ProjectStage
}
