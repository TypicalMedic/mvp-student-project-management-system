package managingstudentproject

type Repository struct {
	id       uint
	name     string
	owner    Account
	branches []Branch
}

func (r *Repository) Branches() []Branch {
	return r.branches
}

func (r *Repository) SetBranches(branches []Branch) {
	r.branches = branches
}

func (r *Repository) Id() uint {
	return r.id
}

func (r *Repository) SetId(id uint) {
	r.id = id
}

func (r *Repository) Name() string {
	return r.name
}

func (r *Repository) Owner() Account {
	return r.owner
}

func (r *Repository) SetName(name string) {
	r.name = name
}

func (r *Repository) SetOwner(owner Account) {
	r.owner = owner
}

func InitRepo(id uint, name string, branches []Branch, owner Account) Repository {
	return Repository{id: id, name: name, owner: owner, branches: branches}
}

// type Repository struct {
// 	name     string
// 	branches []Branch
// 	owner    RepoMngAccount
// }

// func InitRepo(name string, branches []Branch, owner RepoMngAccount) Repository {
// 	return Repository{name: name, branches: branches, owner: owner}
// }

// func (r *Repository) Branches() []Branch {
// 	return r.branches
// }
