package managingstudentproject

type Repository struct {
	name  string
	owner Account
}

func (r *Repository) GetName() string {
	return r.name
}

func (r *Repository) GetOwner() Account {
	return r.owner
}

func (r *Repository) SetName(name string) {
	r.name = name
}

func (r *Repository) SetOwner(owner Account) {
	r.owner = owner
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
