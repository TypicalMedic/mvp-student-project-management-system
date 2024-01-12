package managingstudentproject

type Repository struct {
	name  string
	owner Account
}

func InitRepo(name string, branches []Branch, owner Account) Repository {
	return Repository{name: name, owner: owner}
}

// type Repository struct {
// 	name     string
// 	branches []Branch
// 	owner    RepoMngAccount
// }

// func InitRepo(name string, branches []Branch, owner RepoMngAccount) Repository {
// 	return Repository{name: name, branches: branches, owner: owner}
// }

// func (r *Repository) GetBranches() []Branch {
// 	return r.branches
// }
