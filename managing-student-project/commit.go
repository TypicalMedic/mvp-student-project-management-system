package managingstudentproject

import "time"

type Commit struct {
	id       string
	commiter Account
	message  string
	date     time.Time
	branch   Branch
}

func InitCommit(id string, msg string, commiter Account, date time.Time, branch Branch) Commit {
	return Commit{id: id, message: msg, commiter: commiter, date: date, branch: branch}
}
