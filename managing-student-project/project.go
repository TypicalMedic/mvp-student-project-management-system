package managingstudentproject

import "github.com/google/uuid"

type ProjectStatus int

const (
	ProjectNotConfirmed ProjectStatus = iota
	ProjectInProgress
	ProjectFinished
	ProjectCancelled
)

type Project struct {
	id          uuid.UUID
	theme       string
	timePeriod  string
	supervisor  Professor
	student     Student
	tasks       []Task
	status      ProjectStatus
	repository  Repository
	driveFolder Folder
}

// add lower interface
// func (p *Project) GiveTask(task Task) {
// 	p.tasks = append(p.tasks, task)
// }

func InitProject(theme string, student Student, status ProjectStatus, timePeriod string, repo Repository, tasks []Task, supervisor Professor) Project {
	return Project{id: uuid.New(), theme: theme, student: student, tasks: tasks, status: status, timePeriod: timePeriod, repository: repo, supervisor: supervisor}
}

func (p *Project) GetId() uuid.UUID {
	return p.id
}

func (p *Project) GetRepo() Repository {
	return p.repository
}

func (p *Project) ConnectDriveFolder(folder Folder) {
	p.driveFolder = folder
}

func (p *Project) GetDriveFolder() Folder {
	return p.driveFolder
}
