package models

import "time"

type Professor struct {
	Id                int
	Name              string
	Surname           string
	MiddleName        string
	Position          string
	CalendarAccountId int
	CalendarId        string
	RepoAccountId     int
}

type Project struct {
	Id                 int
	ThemeId            int
	FromYear           int
	ToYear             int
	SupervisorId       int
	StudentId          int
	StatusId           int
	RepoId             int
	CloudDriveFolderId int
}

type Student struct {
	Id         int
	Name       string
	Surname    string
	MiddleName string
	GroupId    int
}

type Task struct {
	Id                 int
	Name               string
	Description        string
	Deadline           time.Time
	StatusId           int
	ProjectStageId     int
	CloudDriveFolderId int
}

type Theme struct {
	Id   int
	Name string
}

type ProjectStatus struct {
	Id   int
	Name string
}

type ProjectStage struct {
	Id   int
	Name string
}

type TaskStatus struct {
	Id   int
	Name string
}

type MeetingStatus struct {
	Id   int
	Name string
}

type Repository struct {
	Id    int
	Name  string
	Owner string
}

type CalendarAccount struct {
	Id    int
	Login string
	Email string
}

type RepositoryAccount struct {
	Id    int
	Login string
	Email string
}

type UniGroup struct {
	Id        int
	CourceId  int
	StartYear int
}

type Cource struct {
	Id   int
	Name string
}

type Meeting struct {
	Id            int
	OrganizerId   int
	ParticipantId int
	Time          time.Time
	IsOnline      bool
	StatusId      int
	Description   string
}
