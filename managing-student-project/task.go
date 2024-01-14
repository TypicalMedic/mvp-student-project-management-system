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
	id           int
	name         string
	deadline     time.Time
	description  string
	status       TaskStatus
	projectStage ProjectStage
	taskFolder   Folder
}

func (t *Task) GetId() int {
	return t.id
}

func (t *Task) GetName() string {
	return t.name
}

func (t *Task) GetDeadline() time.Time {
	return t.deadline
}

func (t *Task) GetDescription() string {
	return t.description
}

func (t *Task) GetStatus() TaskStatus {
	return t.status
}

func (t *Task) GetProjectStage() ProjectStage {
	return t.projectStage
}

func (t *Task) GetTaskFolder() Folder {
	return t.taskFolder
}

func (t *Task) SetId(id int) {
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

func InitTask(name string, deadline time.Time, desc string, status TaskStatus, pStage ProjectStage) Task {
	return Task{name: name, deadline: deadline, description: desc, status: status, projectStage: pStage}
}

func (t *Task) ConnectDriveFolder(folder Folder) {
	t.taskFolder = folder
}

func (t *Task) GetDriveFolder() Folder {
	return t.taskFolder
}
