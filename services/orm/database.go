package orm

import (
	"mvp-student-project-management-system/services/orm/dbinteractors"

	"gorm.io/gen"
	"gorm.io/gorm"
)

type Database struct {
	dbInteractor *dbinteractors.Query
}

func InitDatabade(db *gorm.DB, opts ...gen.DOOption) Database {
	return Database{dbInteractor: dbinteractors.Use(db, opts...)}
}
