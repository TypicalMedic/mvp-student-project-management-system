package managingstudentproject

import (
	"fmt"
	"time"
)

type TaskStatus int

const (
	NotStarted TaskStatus = iota
	InProgress
	Finished
)

func (s TaskStatus) String() string {
	switch s {
	case TaskStatus(NotStarted):
		return "NotStarted"
	case TaskStatus(InProgress):
		return "InProgress"
	case TaskStatus(Finished):
		return "Finished"
	default:
		return fmt.Sprintf("%d", int(s))
	}
}

type Task struct {
	id           uint
	name         string
	deadline     time.Time
	description  string
	status       TaskStatus
	projectStage ProjectStage
	taskFolder   Folder
}

func (t *Task) Id() uint {
	return t.id
}

func (t *Task) Name() string {
	return t.name
}

func (t *Task) Deadline() time.Time {
	return t.deadline
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) Status() TaskStatus {
	return t.status
}

func (t *Task) ProjectStage() ProjectStage {
	return t.projectStage
}

func (t *Task) TaskFolder() Folder {
	return t.taskFolder
}

func (t *Task) SetId(id uint) {
	t.id = id
}

func (t *Task) SetName(name string) {
	t.name = name
}

func (t *Task) SetDeadline(deadline time.Time) {
	t.deadline = deadline
}

func (t *Task) SetDescription(description string) {
	t.description = description
}

func (t *Task) SetStatus(status TaskStatus) {
	t.status = status
}

func (t *Task) SetProjectStage(projectStage ProjectStage) {
	t.projectStage = projectStage
}

func (t *Task) SetTaskFolder(taskFolder Folder) {
	t.taskFolder = taskFolder
}

func InitTask(id uint, name string, deadline time.Time, desc string, status TaskStatus, pStage ProjectStage) Task {
	return Task{id: id, name: name, deadline: deadline, description: desc, status: status, projectStage: pStage}
}

func (t *Task) ConnectDriveFolder(folder Folder) {
	t.taskFolder = folder
}
