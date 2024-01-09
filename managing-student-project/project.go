package managingstudentproject

type ProjectStatus int

const (
	ProjectNotConfirmed ProjectStatus = iota
	ProjectInProgress
	ProjectFinished
	ProjectCancelled
)

type Project struct {
	theme   string
	student Student
	tasks   []Task
	status  ProjectStatus
}

// add lower interface
func (p *Project) GiveTask(task Task) {
	p.tasks = append(p.tasks, task)
}
