package todatabase

import (
	entities "mvp-student-project-management-system/managing-student-project"
	model "mvp-student-project-management-system/services/database/models"
)

func CalendarAccount(account entities.Account) model.CalendarAccount {
	return model.CalendarAccount{
		ID:    int32(account.Id()),
		Login: account.AccountName(),
		Email: account.Email(),
	}
}

func MeetingStatus(status entities.MeetingStatus) model.MeetingStatus {
	return model.MeetingStatus{
		ID:   int32(status),
		Name: status.String(),
	}
}

func Meeting(meeting entities.Meeting) model.Meeting {
	org := meeting.Organizer()
	prt := meeting.Participant()
	return model.Meeting{
		ID:            int32(meeting.Id()),
		OrganizerID:   int32(org.Id()),
		ParticipantID: int32(prt.Id()),
		Time:          meeting.Time(),
		IsOnline:      meeting.IsOnline(),
		StatusID:      int32(meeting.Status()),
		Description:   meeting.Description(),
	}
}

func Professor(professor entities.Professor) model.Professor {
	fullname := professor.FullName()
	calendar := professor.Calendar()
	cowner := calendar.Owner()
	repoacc := professor.RepoMngAccount()
	return model.Professor{
		ID:                int32(professor.Id()),
		Name:              fullname.Name(),
		Surname:           fullname.Surname(),
		MiddleName:        fullname.MiddleName(),
		Position:          professor.Position(),
		CalendarAccountID: int32(cowner.Id()),
		CalendarID:        calendar.Id(),
		RepoAccountID:     int32(repoacc.Id()),
	}
}

func ProjectStage(stage entities.ProjectStage) model.ProjectStage {
	return model.ProjectStage{
		ID:   int32(stage),
		Name: stage.String(),
	}
}

func ProjectStatus(status entities.ProjectStatus) model.ProjectStatus {
	return model.ProjectStatus{
		ID:   int32(status),
		Name: status.String(),
	}
}

func Project(projcet entities.Project) model.Project {
	sup := projcet.Supervisor()
	stud := projcet.Student()
	repo := projcet.Repository()
	cloudf := projcet.DriveFolder()
	return model.Project{
		ID:                 int32(projcet.Id()),
		Theme:              projcet.Theme(),
		Year:               int32(projcet.Year()),
		SupervisorID:       int32(sup.Id()),
		StudentID:          int32(stud.Id()),
		StatusID:           int32(projcet.Status()),
		RepoID:             int32(repo.Id()),
		CloudDriveFolderID: cloudf.Id(),
	}
}

func RepositoryAccount(account entities.Account) model.RepositoryAccount {
	return model.RepositoryAccount{
		ID:    int32(account.Id()),
		Login: account.AccountName(),
		Email: account.Email(),
	}
}

func Repository(repository entities.Repository) model.Repository {
	owner := repository.Owner()
	return model.Repository{
		ID:    int32(repository.Id()),
		Name:  repository.Name(),
		Owner: owner.AccountName(),
	}
}

func Student(student entities.Student) model.Student {
	fullname := student.FullName()
	return model.Student{
		ID:         int32(student.Id()),
		Name:       fullname.Name(),
		Surname:    fullname.Surname(),
		MiddleName: fullname.MiddleName(),
		UniGroup:   student.Group(),
	}
}

func TaskStatus(status entities.TaskStatus) model.TaskStatus {
	return model.TaskStatus{
		ID:   int32(status),
		Name: status.String(),
	}
}

func Task(task entities.Task) model.Task {
	cloudf := task.TaskFolder()
	return model.Task{
		ID:                 int32(task.Id()),
		Name:               task.Name(),
		Description:        task.Description(),
		Deadline:           task.Deadline(),
		StatusID:           int32(task.Status()),
		ProjectStageID:     int32(task.ProjectStage()),
		CloudDriveFolderID: cloudf.Id(),
	}
}
