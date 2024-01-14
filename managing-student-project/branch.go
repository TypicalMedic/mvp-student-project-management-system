package managingstudentproject

type Branch struct {
	name       string
	repository Repository
}

func (b *Branch) GetName() string {
	return b.name
}

func (b *Branch) GetRepository() Repository {
	return b.repository
}

func (b *Branch) SetName(name string) {
	b.name = name
}

func (b *Branch) SetRepository(repository Repository) {
	b.repository = repository
}

func InitBranch(name string, repo Repository) Branch {
	return Branch{name: name, repository: repo}
}

// type Branch struct {
// 	name    string
// 	commits []Commit //?????????????
// }

// func InitBranch(name string, commits []Commit) Branch {
// 	return Branch{name: name, commits: commits}
// }

// func (b *Branch) GetCommits() []Commit {
// 	return b.commits
// }
