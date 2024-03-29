package usecaselayer

import (
	entities "mvp-student-project-management-system/managing-student-project"
	interfaces "mvp-student-project-management-system/use-case-layer/interfaces"
	"time"

	irepositories "mvp-student-project-management-system/use-case-layer/interfaces/irepositories"
)

type ProjectInteractor struct {
	ProjectDataStorage irepositories.IProjectRepository
	Presenter          interfaces.IPresenter
	RepoManager        interfaces.IRepoManager
	Drive              interfaces.ICloudDrive
}

func (p *ProjectInteractor) ViewProfProjectsList(professor entities.Professor) {
	p.Presenter.ShowProjectsList(p.ProjectDataStorage.GetProfessorProjects(professor))
}

func (p *ProjectInteractor) ViewProject(project entities.Project) {
	p.Presenter.ShowProject(project)
}

func (p *ProjectInteractor) ViewRepoCommitsFromTime(repo entities.Repository, fromTime time.Time) {
	branches := p.RepoManager.GetRepoBranches(repo)
	commits := []entities.Commit{}
	for _, branch := range branches {
		bCommits := p.RepoManager.GetBranchCommitsFromTime(branch, fromTime)
		commits = append(commits, bCommits...)
	}
	p.Presenter.ShowCommits(commits)
}

func (p *ProjectInteractor) CreateProject(project entities.Project) {
	projectFolder := p.Drive.CteateProjectFolder(project)
	project.ConnectDriveFolder(projectFolder)
	p.ProjectDataStorage.SaveProject(project)
}

func (p *ProjectInteractor) DeleteProject(project entities.Project) {
	p.Drive.DeleteFolder(project.DriveFolder())
	p.ProjectDataStorage.DeleteProject(project)
}

func (p *ProjectInteractor) SelectProjectFromList(list []entities.Project, projectId uint) entities.Project {
	for _, project := range list {
		if project.Id() == projectId {
			return project
		}
	}
	return entities.Project{}
}
