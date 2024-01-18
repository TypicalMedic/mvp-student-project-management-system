package github

import (
	"context"
	"time"

	"github.com/google/go-github/v56/github"
)

// думаю есть возможность сделать как у гугла генерацию токена
const AUTH_TOKEN = "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"

type githubAPI struct {
	Context context.Context
	Client  *github.Client
}

func InitGithubAPI() githubAPI {
	context := context.Background()
	client := github.NewClient(nil).WithAuthToken(AUTH_TOKEN)

	return githubAPI{Context: context, Client: client}
}

func (g *githubAPI) GetRepos() ([]*github.Repository, error) {
	reps, _, err := g.Client.Repositories.List(g.Context, "", nil)
	if err == nil {
		return reps, nil
	}
	return nil, err
}

func (g *githubAPI) GetRepoCommitsFromTime(owner, repoName string, fromTime time.Time) ([]*github.RepositoryCommit, error) {
	opt := &github.CommitsListOptions{Since: fromTime}
	commits, _, err := g.Client.Repositories.ListCommits(g.Context, owner, repoName, opt)
	if err == nil {
		return commits, nil
	}
	return nil, err
}

func (g *githubAPI) GetBranchCommitsFromTime(owner, repoName string, branch string, fromTime time.Time) ([]*github.RepositoryCommit, error) {
	opt := &github.CommitsListOptions{SHA: branch, Since: fromTime}
	commits, _, err := g.Client.Repositories.ListCommits(g.Context, owner, repoName, opt)
	if err == nil {
		return commits, nil
	}
	return nil, err
}

func (g *githubAPI) GetRepoBranches(owner, repoName string) ([]*github.Branch, error) {
	opt := &github.BranchListOptions{}
	branches, _, err := g.Client.Repositories.ListBranches(g.Context, owner, repoName, opt)
	if err == nil {
		return branches, nil
	}
	return nil, err
}
