package database

import (
	"fmt"
	dbmodels "mvp-student-project-management-system/services/database/models"

	"github.com/huandu/go-sqlbuilder"
)

func SQLInsertProject(project dbmodels.Project) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("project").Cols(
		"id",
		"theme_id",
		"from_year",
		"to_year",
		"supervisor_id",
		"student_id",
		"status_id",
		"repoId",
		"cloud_drive_folder_id").Values(
		project.Id,
		project.ThemeId,
		project.FromYear,
		project.ToYear,
		project.SupervisorId,
		project.StudentId,
		project.StatusId,
		project.RepoId,
		project.CloudDriveFolderId)
	return insert.Build()
}

func SQLInsertCalendarAccount(calendaraccount dbmodels.CalendarAccount) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("calendar_account").Cols(
		"id",
		"login",
		"email").Values(
		calendaraccount.Id,
		calendaraccount.Login,
		calendaraccount.Email)
	return insert.Build()
}

func SQLInsertCource(cource dbmodels.Cource) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("cource").Cols(
		"id",
		"name").Values(
		cource.Id,
		cource.Name)
	return insert.Build()
}
func SQLInsertProfessor(professor dbmodels.Professor) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("professor").Cols(
		"id",
		"name",
		"surname",
		"middle_name",
		"position",
		"calendar_account_id",
		"calendar_id",
		"repo_account_id").Values(
		professor.Id,
		professor.Name,
		professor.Surname,
		professor.MiddleName,
		professor.Position,
		professor.CalendarAccountId,
		professor.CalendarId,
		professor.RepoAccountId)
	return insert.Build()
}

func SQLInsertProjectStage(projectstage dbmodels.ProjectStage) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("project_stage").Cols(
		"id",
		"name").Values(
		projectstage.Id,
		projectstage.Name)
	return insert.Build()
}
func SQLInsertProjectStatus(projectstatus dbmodels.ProjectStatus) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("project_status").Cols(
		"id",
		"name").Values(
		projectstatus.Id,
		projectstatus.Name)
	return insert.Build()
}

func SQLInsertRepository(repository dbmodels.Repository) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("repository").Cols(
		"id",
		"name",
		"owner").Values(
		repository.Id,
		repository.Name,
		repository.Owner)
	return insert.Build()
}

func SQLInsertRepositoryAccount(repositoryaccount dbmodels.RepositoryAccount) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("repository_account").Cols(
		"id",
		"login",
		"email").Values(
		repositoryaccount.Id,
		repositoryaccount.Login,
		repositoryaccount.Email)
	return insert.Build()
}

func SQLInsertStudent(student dbmodels.Student) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("student").Cols(
		"id",
		"name",
		"surname",
		"middle_name",
		"group_id").Values(
		student.Id,
		student.Name,
		student.Surname,
		student.MiddleName,
		student.GroupId)
	return insert.Build()
}

func SQLInsertTask(task dbmodels.Task) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("task").Cols(
		"id",
		"name",
		"description",
		"deadline",
		"status_id",
		"project_stage_id",
		"cloud_drive_folder_id").Values(
		task.Id,
		task.Name,
		task.Description,
		task.Deadline,
		task.StatusId,
		task.ProjectStageId,
		task.CloudDriveFolderId)
	return insert.Build()
}

func SQLInsertTaskStatus(taskstatus dbmodels.TaskStatus) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("task_status").Cols(
		"id",
		"name").Values(
		taskstatus.Id,
		taskstatus.Name)
	return insert.Build()
}

func SQLInsertTheme(theme dbmodels.Theme) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("theme").Cols(
		"id",
		"name").Values(
		theme.Id,
		theme.Name)
	return insert.Build()
}

func SQLInsertUniGroup(unigroup dbmodels.UniGroup) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("uni_group").Cols(
		"id",
		"cource_id",
		"start_year").Values(
		unigroup.Id,
		unigroup.CourceId,
		unigroup.StartYear)
	return insert.Build()
}

func SQLInsertMeeting(meeting dbmodels.Meeting) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("meeting").Cols(
		"id",
		"organizer_id",
		"participant_id",
		"time",
		"is_online",
		"status_id",
		"description").Values(
		meeting.Id,
		meeting.OrganizerId,
		meeting.ParticipantId,
		meeting.Time,
		meeting.IsOnline,
		meeting.StatusId,
		meeting.Description)
	return insert.Build()
}

func SQLInsertMeetingStatus(meetingstatus dbmodels.MeetingStatus) (string, []interface{}) {
	insert := sqlbuilder.NewInsertBuilder()
	insert.InsertInto("meeting_status").Cols(
		"id",
		"name").Values(
		meetingstatus.Id,
		meetingstatus.Name)
	return insert.Build()
}

func SQLDeleteProject(project dbmodels.Project) (string, []interface{}) {
	delete := sqlbuilder.NewDeleteBuilder()
	delete.DeleteFrom("project").Where(fmt.Sprint("id = ", project.Id, ";"))
	return delete.Build()
}

func SQLDeleteMeeting(meeting dbmodels.Meeting) (string, []interface{}) {
	delete := sqlbuilder.NewDeleteBuilder()
	delete.DeleteFrom("meeting").Where(fmt.Sprint("id = ", meeting.Id, ";"))
	return delete.Build()
}

func SQLDeleteTask(task dbmodels.Task) (string, []interface{}) {
	delete := sqlbuilder.NewDeleteBuilder()
	delete.DeleteFrom("task").Where(fmt.Sprint("id = ", task.Id, ";"))
	return delete.Build()
}
