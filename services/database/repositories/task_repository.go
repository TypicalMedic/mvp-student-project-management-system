package repositories

import (
	"log"
	entities "mvp-student-project-management-system/managing-student-project"
	"mvp-student-project-management-system/services/database"
	"mvp-student-project-management-system/services/database/datagateway/todatabase"
	"mvp-student-project-management-system/services/database/datagateway/todomain"
)

type DBTaskRepository struct {
	DBContext database.Database
}

func (r *DBTaskRepository) SaveTask(task entities.Task) {
	dbtask := todatabase.Task(task)
	err := r.DBContext.DBInteractor.Task.Create(&dbtask)
	if err != nil {
		log.Fatal(err)
	}
}
func (r *DBTaskRepository) DeleteTask(task entities.Task) {
	dbtask := todatabase.Task(task)
	_, err := r.DBContext.DBInteractor.Task.Delete(&dbtask)
	if err != nil {
		log.Fatal(err)
	}
}
func (r *DBTaskRepository) UpdateTask(task entities.Task) {
	dbtask := todatabase.Task(task)
	err := r.DBContext.DBInteractor.Task.Save(&dbtask)
	if err != nil {
		log.Fatal(err)
	}
}
func (r *DBTaskRepository) GetProjectTasks(project entities.Project) []entities.Task {
	t := r.DBContext.DBInteractor.Task
	dbproject := todatabase.Project(project)
	tasks, err := t.Where(t.ProjectID.Eq(dbproject.ID)).Find()
	if err != nil {
		log.Fatal(err)
	}
	out := []entities.Task{}
	for _, task := range tasks {
		out = append(out, todomain.Task(*task))
	}
	return out
}
