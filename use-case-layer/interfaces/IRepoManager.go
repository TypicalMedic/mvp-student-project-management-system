package interfaces

import (
	entities "mvp-student-project-management-system/managing-student-project"
	"time"
)

type IRepoManager interface {
	GetRepoBranches(entities.Repository) []entities.Branch
	GetBranchCommitsFromTime(entities.Branch, time.Time) []entities.Commit
}
