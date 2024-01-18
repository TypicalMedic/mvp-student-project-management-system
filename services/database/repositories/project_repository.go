package repositories

import (
	"log"
	entities "mvp-student-project-management-system/managing-student-project"
	"mvp-student-project-management-system/services/database"
	"mvp-student-project-management-system/services/database/datagateway/todatabase"
	"mvp-student-project-management-system/services/database/datagateway/todomain"
)

type DBProjectRepository struct {
	DBContext database.Database
}

func (r *DBProjectRepository) SaveProject(project entities.Project) {
	dbproject := todatabase.Project(project)
	err := r.DBContext.DBInteractor.Project.Create(&dbproject)
	if err != nil {
		log.Fatal(err)
	}
}
func (r *DBProjectRepository) DeleteProject(project entities.Project) {
	dbproject := todatabase.Project(project)
	_, err := r.DBContext.DBInteractor.Project.Delete(&dbproject)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *DBProjectRepository) GetProfessorProjects(professor entities.Professor) []entities.Project {
	p := r.DBContext.DBInteractor.Project
	dbprofessor := todatabase.Professor(professor)
	projects, err := p.Where(p.SupervisorID.Eq(dbprofessor.ID)).Find()
	if err != nil {
		log.Fatal(err)
	}
	out := []entities.Project{}
	for _, project := range projects {
		out = append(out, todomain.Project(*project))
	}
	return out
}
