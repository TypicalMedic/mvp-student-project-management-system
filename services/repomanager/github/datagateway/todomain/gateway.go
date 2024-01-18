package todomain

import (
	entities "mvp-student-project-management-system/managing-student-project"

	"github.com/google/go-github/v56/github"
)

func Repository(repository github.Repository, setters ...func(*entities.Repository)) entities.Repository {
	r := entities.InitRepo(
		uint(*repository.ID),
		*repository.Name,
		[]entities.Branch{},
		entities.InitAccount(0, *repository.Owner.Name, *repository.Owner.Email))
	for _, setter := range setters {
		setter(&r)
	}
	return r
}

func Branch(branch github.Branch, setters ...func(*entities.Branch)) entities.Branch {
	b := entities.InitBranch(
		*branch.Name,
		entities.InitRepo(0, "", []entities.Branch{}, entities.Account{}),
		[]entities.Commit{})
	for _, setter := range setters {
		setter(&b)
	}
	return b
}

func Commit(commit github.Commit, setters ...func(*entities.Commit)) entities.Commit {
	c := entities.InitCommit(
		*commit.SHA,
		*commit.Message,
		entities.InitAccount(0, *commit.Committer.Name, *commit.Committer.Email),
		commit.Committer.Date.Time)
	for _, setter := range setters {
		setter(&c)
	}
	return c
}
