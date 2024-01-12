package usecaselayer

import (
	entities "mvp-student-project-management-system/managing-student-project"
	interfaces "mvp-student-project-management-system/use-case-layer/interfaces"
)

type TaskInteractor struct {
	TaskDataStorage interfaces.ITaskDataStorage
	Presenter       interfaces.IPresenter
	Drive           interfaces.ICloudDrive
}

func (t *TaskInteractor) ViewProjectTasksList(project entities.Project) {
	tasks := t.TaskDataStorage.GetProjectTasks(project)
	t.Presenter.ShowTasks(tasks)
}

func (t *TaskInteractor) GiveProjectTask(task entities.Task, projectFolder entities.Folder) {
	tFolder := t.Drive.CteateTaskFolder(task, projectFolder)
	t.Drive.CreateTaskFile(tFolder, task)
	task.ConnectDriveFolder(tFolder)
	t.TaskDataStorage.SaveTask(task)
}

func (t *TaskInteractor) DeleteTask(task entities.Task) {
	t.Drive.DeleteFolder(task.GetDriveFolder())
	t.TaskDataStorage.DeleteTask(task)
}
