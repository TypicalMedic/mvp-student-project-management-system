package commit

import (
	entities "mvp-student-project-management-system/managing-student-project"
)

func SetCommiter(commiter entities.Account) func(*entities.Commit) {
	return func(c *entities.Commit) {
		c.SetCommiter(commiter)
	}
}
