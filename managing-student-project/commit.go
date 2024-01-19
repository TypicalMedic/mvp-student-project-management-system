package managingstudentproject

import "time"

type Commit struct {
	id       string
	commiter Account
	message  string
	date     time.Time
}

func (c *Commit) Id() string {
	return c.id
}

func (c *Commit) Commiter() Account {
	return c.commiter
}

func (c *Commit) Message() string {
	return c.message
}

func (c *Commit) Date() time.Time {
	return c.date
}

func (c *Commit) SetId(id string) {
	c.id = id
}

func (c *Commit) SetCommiter(commiter Account) {
	c.commiter = commiter
}

func (c *Commit) SetMessage(message string) {
	c.message = message
}

func (c *Commit) SetDate(date time.Time) {
	c.date = date
}

func InitCommit(id string, msg string, commiter Account, date time.Time) Commit {
	return Commit{id: id, message: msg, commiter: commiter, date: date}
}
