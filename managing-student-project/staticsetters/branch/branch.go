package branch

import (
	entities "mvp-student-project-management-system/managing-student-project"
)

func SetCommits(commits []entities.Commit) func(*entities.Branch) {
	return func(b *entities.Branch) {
		b.SetCommits(commits)
	}
}

func SetRepository(repo entities.Repository) func(*entities.Branch) {
	return func(b *entities.Branch) {
		b.SetRepository(repo)
	}
}
