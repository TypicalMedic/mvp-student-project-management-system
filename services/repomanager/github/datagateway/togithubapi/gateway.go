package togithubapi

import (
	entities "mvp-student-project-management-system/managing-student-project"

	"github.com/google/go-github/v56/github"
)

func Repository(repository entities.Repository) github.Repository {
	owner := repository.Owner()
	r := github.Repository{
		ID:   github.Int64(int64(repository.Id())),
		Name: github.String(repository.Name()),
		Owner: &github.User{
			Name:  github.String(owner.AccountName()),
			Email: github.String(owner.Email()),
		},
	}
	return r
}

func Branch(branch entities.Branch) github.Branch {
	b := github.Branch{
		Name: github.String(branch.Name()),
	}
	return b
}

func Commit(commit entities.Commit) github.Commit {
	commiter := commit.Commiter()
	c := github.Commit{
		SHA:     github.String(commit.Id()),
		Message: github.String(commit.Message()),
		Committer: &github.CommitAuthor{
			Date: &github.Timestamp{
				commit.Date(),
			},
			Name:  github.String(commiter.AccountName()),
			Email: github.String(commiter.Email()),
		},
	}
	return c
}
