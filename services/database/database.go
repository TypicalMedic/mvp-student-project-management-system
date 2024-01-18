package database

import (
	"mvp-student-project-management-system/services/database/dbinteractors"

	"gorm.io/gen"
	"gorm.io/gorm"
)

type Database struct {
	DBInteractor *dbinteractors.Query
}

func InitDatabade(db *gorm.DB, opts ...gen.DOOption) Database {
	return Database{DBInteractor: dbinteractors.Use(db, opts...)}
}
