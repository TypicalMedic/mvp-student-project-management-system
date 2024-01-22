package googledrive

import (
	"fmt"
	"log"
	entities "mvp-student-project-management-system/managing-student-project"
	"mvp-student-project-management-system/services/clouddrive/googledrive/datagateway/todomain"
)

type GoogleDrive struct {
	api googleDriveApi
}

func InitGoogleDrive(api googleDriveApi) GoogleDrive {
	return GoogleDrive{api: api}
}

func (d *GoogleDrive) CteateProjectFolder(project entities.Project) entities.Folder {
	student := project.Student()
	stname := student.FullName()
	sup := project.Supervisor()
	supFolder := sup.DriveFolder()
	gdFolder, err := d.api.CreateFolder(supFolder.Id(), fmt.Sprint(project.Theme(), "_", stname.ToString(), "_", project.Year()))
	if err != nil {
		log.Fatal(err)
	}
	folder := todomain.Folder(*gdFolder)
	return folder
}
func (d *GoogleDrive) CteateTaskFolder(task entities.Task, parentFolder entities.Folder) entities.Folder {
	gdFolder, err := d.api.CreateFolder(parentFolder.Id(), fmt.Sprint("Task_", task.Name(), "_due_to_", task.Deadline()))
	if err != nil {
		log.Fatal(err)
	}
	folder := todomain.Folder(*gdFolder)
	return folder
}
func (d *GoogleDrive) DeleteFolder(folder entities.Folder) {
	err := d.api.DeleteFolder(folder.Id())
	if err != nil {
		log.Fatal(err)
	}
}
func (d *GoogleDrive) CreateTaskFile(folder entities.Folder, task entities.Task) {
	_, err := d.api.AddTextFileToFolder(folder.Id(), "Task description", fmt.Sprint(task.Name(), "\nDescription: \n", task.Description()))
	if err != nil {
		log.Fatal(err)
	}
}
