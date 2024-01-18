package managingstudentproject

type Repository struct {
	id      uint
	name    string
	owner   Account
	commits []Commit
}

func (r *Repository) Commits() []Commit {
	return r.commits
}

func (r *Repository) SetCommits(commits []Commit) {
	r.commits = commits
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
	return Repository{id: id, name: name, owner: owner}
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
