package repository

import (
	entities "mvp-student-project-management-system/managing-student-project"
)

func SetBranches(branches []entities.Branch) func(*entities.Repository) {
	return func(p *entities.Repository) {
		p.SetBranches(branches)
	}
}

func SetOwner(owner entities.Account) func(*entities.Repository) {
	return func(p *entities.Repository) {
		p.SetOwner(owner)
	}
}
