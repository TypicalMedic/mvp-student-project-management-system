package managingstudentproject

type Calendar struct {
	id    string
	owner Account
	name  string
}

func (c *Calendar) Id() string {
	return c.id
}

func (c *Calendar) Owner() Account {
	return c.owner
}

func (c *Calendar) Name() string {
	return c.name
}

func (c *Calendar) SetId(id string) {
	c.id = id
}

func (c *Calendar) SetOwner(owner Account) {
	c.owner = owner
}

func (c *Calendar) SetName(name string) {
	c.name = name
}

func InitCalendar(id string, owner Account, name string) Calendar {
	return Calendar{id: id, owner: owner, name: name}
}
