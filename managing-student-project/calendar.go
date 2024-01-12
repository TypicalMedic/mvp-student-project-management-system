package managingstudentproject

type Calendar struct {
	owner Account
	name  string
}

func InitCalendar(owner Account, name string) Calendar {
	return Calendar{owner: owner, name: name}
}
