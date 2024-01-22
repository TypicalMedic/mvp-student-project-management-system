package todomain

import (
	entities "mvp-student-project-management-system/managing-student-project"
	model "mvp-student-project-management-system/services/database/models"
)

func CalendarAccount(account model.CalendarAccount) entities.Account {
	return entities.InitAccount(
		uint(account.ID),
		account.Login,
		account.Email)
}

func MeetingStatus(status model.MeetingStatus) entities.MeetingStatus {
	return entities.MeetingStatus(status.ID)
}

func Meeting(meeting model.Meeting, setters ...func(*entities.Meeting)) entities.Meeting {
	m := entities.InitMeeting(
		uint(meeting.ID),
		entities.InitStudent(uint(meeting.ParticipantID), entities.FullName{}, ""),
		meeting.Time,
		meeting.IsOnline,
		entities.MeetingStatus(meeting.StatusID),
		entities.InitProfessor(uint(meeting.OrganizerID), entities.FullName{}, "", []entities.Meeting{}, []entities.Project{}, entities.Calendar{}, entities.Account{}, entities.Folder{}),
		meeting.Description)
	for _, setter := range setters {
		setter(&m)
	}
	return m
}

func Professor(professor model.Professor, setters ...func(*entities.Professor)) entities.Professor {
	p := entities.InitProfessor(
		uint(professor.ID),
		entities.InitFullName(professor.Name, professor.Surname, professor.MiddleName),
		professor.Position,
		[]entities.Meeting{},
		[]entities.Project{},
		entities.InitCalendar(professor.CalendarID, entities.Account{}, ""),
		entities.InitAccount(uint(professor.RepoAccountID), "", ""),
		entities.InitFolder(professor.CloudDriveFolderID, "", ""),
	)
	for _, setter := range setters {
		setter(&p)
	}
	return p
}

func ProjectStage(stage model.ProjectStage) entities.ProjectStage {
	return entities.ProjectStage(stage.ID)
}

func ProjectStatus(status model.ProjectStatus) entities.ProjectStatus {
	return entities.ProjectStatus(status.ID)
}

func Project(projcet model.Project, setters ...func(*entities.Project)) entities.Project {
	p := entities.InitProject(
		uint(projcet.ID),
		projcet.Theme,
		entities.InitStudent(uint(projcet.StudentID), entities.FullName{}, ""),
		entities.ProjectStatus(projcet.StatusID),
		int(projcet.Year),
		entities.InitRepo(uint(projcet.RepoID), "", []entities.Branch{}, entities.Account{}),
		[]entities.Task{},
		entities.InitProfessor(uint(projcet.SupervisorID), entities.FullName{}, "", []entities.Meeting{}, []entities.Project{}, entities.Calendar{}, entities.Account{}, entities.Folder{}))
	for _, setter := range setters {
		setter(&p)
	}
	return p
}

func RepositoryAccount(account model.RepositoryAccount) entities.Account {
	return entities.InitAccount(
		uint(account.ID),
		account.Login,
		account.Email)
}

func Repository(repository model.Repository) entities.Repository {
	return entities.InitRepo(
		uint(repository.ID),
		repository.Name,
		[]entities.Branch{},
		entities.InitAccount(0, repository.Owner, ""))
}

func Student(student model.Student) entities.Student {
	return entities.InitStudent(
		uint(student.ID),
		entities.InitFullName(student.Name, student.Surname, student.MiddleName),
		student.UniGroup)
}

func TaskStatus(status model.TaskStatus) entities.TaskStatus {
	return entities.TaskStatus(status.ID)
}

func Task(task model.Task) entities.Task {
	return entities.InitTask(
		uint(task.ID),
		task.Name,
		task.Deadline,
		task.Description,
		entities.TaskStatus(task.StatusID),
		entities.ProjectStage(task.ProjectStageID))
}
