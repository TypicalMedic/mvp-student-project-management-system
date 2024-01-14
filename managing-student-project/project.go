package managingstudentproject

type ProjectStatus int

const (
	ProjectNotConfirmed ProjectStatus = iota
	ProjectInProgress
	ProjectFinished
	ProjectCancelled
)

type Project struct {
	id          int
	theme       string
	timePeriod  string
	supervisor  Professor
	student     Student
	tasks       []Task
	status      ProjectStatus
	repository  Repository
	driveFolder Folder
}

func (p *Project) GetTheme() string {
	return p.theme
}

func (p *Project) GetTimePeriod() string {
	return p.timePeriod
}

func (p *Project) GetSupervisor() Professor {
	return p.supervisor
}

func (p *Project) GetStudent() Student {
	return p.student
}

func (p *Project) GetTasks() []Task {
	return p.tasks
}

func (p *Project) GetStatus() ProjectStatus {
	return p.status
}

func (p *Project) GetRepository() Repository {
	return p.repository
}

func (p *Project) GetDriveFolder() Folder {
	return p.driveFolder
}

func (p *Project) SetTheme(theme string) {
	p.theme = theme
}

func (p *Project) SetTimePeriod(timePeriod string) {
	p.timePeriod = timePeriod
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

func InitProject(theme string, student Student, status ProjectStatus, timePeriod string, repo Repository, tasks []Task, supervisor Professor) Project {
	return Project{theme: theme, student: student, tasks: tasks, status: status, timePeriod: timePeriod, repository: repo, supervisor: supervisor}
}

func (p *Project) GetId() int {
	return p.id
}

func (p *Project) SetId(id int) {
	p.id = id
}

func (p *Project) ConnectDriveFolder(folder Folder) {
	p.driveFolder = folder
}
