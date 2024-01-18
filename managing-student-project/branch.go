package managingstudentproject

type Branch struct {
	name       string
	repository Repository
	commits    []Commit
}

func (b *Branch) Commits() []Commit {
	return b.commits
}

func (b *Branch) SetCommits(commits []Commit) {
	b.commits = commits
}

func (b *Branch) Name() string {
	return b.name
}

func (b *Branch) Repository() Repository {
	return b.repository
}

func (b *Branch) SetName(name string) {
	b.name = name
}

func (b *Branch) SetRepository(repository Repository) {
	b.repository = repository
}

func InitBranch(name string, repo Repository, commits []Commit) Branch {
	return Branch{name: name, repository: repo, commits: commits}
}

// type Branch struct {
// 	name    string
// 	commits []Commit //?????????????
// }

// func InitBranch(name string, commits []Commit) Branch {
// 	return Branch{name: name, commits: commits}
// }

// func (b *Branch) Commits() []Commit {
// 	return b.commits
// }
