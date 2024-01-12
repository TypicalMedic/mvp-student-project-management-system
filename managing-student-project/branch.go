package managingstudentproject

type Branch struct {
	name       string
	repository Repository
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
