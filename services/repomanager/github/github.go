package github

import (
	"log"
	entities "mvp-student-project-management-system/managing-student-project"
	"mvp-student-project-management-system/services/repomanager/github/datagateway/todomain"
	"time"
)

type Github struct {
	api githubAPI
}

func (g *Github) GetRepoBranches(repo entities.Repository) []entities.Branch {
	owner := repo.Owner()
	ghbranches, err := g.api.GetRepoBranches(owner.AccountName(), repo.Name())
	if err != nil {
		log.Fatal(err)
	}
	branches := []entities.Branch{}
	for _, ghbranch := range ghbranches {
		br := todomain.Branch(*ghbranch)
		branches = append(branches, br)
	}
	return branches
}
func (g *Github) GetBranchCommitsFromTime(branch entities.Branch, time time.Time) []entities.Commit {
	repo := branch.Repository()
	repoOwner := repo.Owner()
	ghcommits, err := g.api.GetBranchCommitsFromTime(repoOwner.AccountName(), repo.Name(), branch.Name(), time)
	if err != nil {
		log.Fatal(err)
	}
	commits := []entities.Commit{}
	for _, ghcommit := range ghcommits {
		cm := todomain.Commit(*ghcommit.Commit)
		commits = append(commits, cm)
	}
	return commits
}
