package database

import (
	"database/sql"
	"log"
	entities "mvp-student-project-management-system/managing-student-project"
)

type Database struct {
	db *sql.DB
}

func InitDatabade(db *sql.DB) Database {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return Database{db: db}
}
func (db *Database) Ping() {
	if err := db.db.Ping(); err != nil {
		log.Fatal(err)
	}
}

// project
func (db *Database) SaveProject(project entities.Project) {
}
func (db *Database) DeleteProject(entities.Project) {}
func (db *Database) GetProfessorProjects(entities.Professor) []entities.Project {
	return []entities.Project{}
}

// meeting
func (db *Database) SaveMeeting(entities.Meeting)   {}
func (db *Database) DeleteMeeting(entities.Meeting) {}
func (db *Database) UpdateMeeting(entities.Meeting) {}

// task
func (db *Database) SaveTask(entities.Task)                           {}
func (db *Database) DeleteTask(entities.Task)                         {}
func (db *Database) UpdateTask(entities.Task)                         {}
func (db *Database) GetProjectTasks(entities.Project) []entities.Task { return []entities.Task{} }
