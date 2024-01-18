package managingstudentproject

import "fmt"

type ProjectStatus int

const (
	ProjectNotConfirmed ProjectStatus = iota
	ProjectInProgress
	ProjectFinished
	ProjectCancelled
)

func (s ProjectStatus) String() string {
	switch s {
	case ProjectStatus(ProjectNotConfirmed):
		return "NotConfirmed"
	case ProjectStatus(ProjectInProgress):
		return "InProgress"
	case ProjectStatus(ProjectFinished):
		return "Finished"
	case ProjectStatus(ProjectCancelled):
		return "Cancelled"
	default:
		return fmt.Sprintf("%d", int(s))
	}
}

type ProjectStage int

const (
	Analysis ProjectStage = iota
	Design
	Development
	Testing
	Deployment
)

func (s ProjectStage) String() string {
	switch s {
	case ProjectStage(Analysis):
		return "Analysis"
	case ProjectStage(Design):
		return "Design"
	case ProjectStage(Development):
		return "Development"
	case ProjectStage(Testing):
		return "Testing"
	case ProjectStage(Deployment):
		return "Deployment"
	default:
		return fmt.Sprintf("%d", int(s))
	}
}

type Project struct {
	id          uint
	theme       string
	year        int
	supervisor  Professor
	student     Student
	tasks       []Task
	status      ProjectStatus
	repository  Repository
	driveFolder Folder
}

func (p *Project) Year() int {
	return p.year
}

func (p *Project) SetYear(from int) {
	p.year = from
}

func (p *Project) Theme() string {
	return p.theme
}

func (p *Project) Supervisor() Professor {
	return p.supervisor
}

func (p *Project) Student() Student {
	return p.student
}

func (p *Project) Tasks() []Task {
	return p.tasks
}

func (p *Project) Status() ProjectStatus {
	return p.status
}

func (p *Project) Repository() Repository {
	return p.repository
}

func (p *Project) DriveFolder() Folder {
	return p.driveFolder
}

func (p *Project) SetTheme(theme string) {
	p.theme = theme
}

func (p *Project) SetSupervisor(supervisor Professor) {
	p.supervisor = supervisor
}

func (p *Project) SetStudent(student Student) {
	p.student = student
}

func (p *Project) SetTasks(tasks []Task) {
	p.tasks = tasks
}

func (p *Project) SetStatus(status ProjectStatus) {
	p.status = status
}

func (p *Project) SetRepository(repository Repository) {
	p.repository = repository
}

func (p *Project) SetDriveFolder(driveFolder Folder) {
	p.driveFolder = driveFolder
}

// add lower interface
// func (p *Project) GiveTask(task Task) {
// 	p.tasks = append(p.tasks, task)
// }

func InitProject(id uint, theme string, student Student, status ProjectStatus, from int, repo Repository, tasks []Task, supervisor Professor) Project {
	return Project{id: id, theme: theme, student: student, tasks: tasks, status: status, year: from, repository: repo, supervisor: supervisor}
}

func (p *Project) Id() uint {
	return p.id
}

func (p *Project) SetId(id uint) {
	p.id = id
}

func (p *Project) ConnectDriveFolder(folder Folder) {
	p.driveFolder = folder
}
