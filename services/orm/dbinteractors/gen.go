// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dbinteractors

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q                 = new(Query)
	CalendarAccount   *calendarAccount
	Meeting           *meeting
	MeetingStatus     *meetingStatus
	Professor         *professor
	Project           *project
	ProjectStage      *projectStage
	ProjectStatus     *projectStatus
	Repository        *repository
	RepositoryAccount *repositoryAccount
	Student           *student
	Task              *task
	TaskStatus        *taskStatus
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	CalendarAccount = &Q.CalendarAccount
	Meeting = &Q.Meeting
	MeetingStatus = &Q.MeetingStatus
	Professor = &Q.Professor
	Project = &Q.Project
	ProjectStage = &Q.ProjectStage
	ProjectStatus = &Q.ProjectStatus
	Repository = &Q.Repository
	RepositoryAccount = &Q.RepositoryAccount
	Student = &Q.Student
	Task = &Q.Task
	TaskStatus = &Q.TaskStatus
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                db,
		CalendarAccount:   newCalendarAccount(db, opts...),
		Meeting:           newMeeting(db, opts...),
		MeetingStatus:     newMeetingStatus(db, opts...),
		Professor:         newProfessor(db, opts...),
		Project:           newProject(db, opts...),
		ProjectStage:      newProjectStage(db, opts...),
		ProjectStatus:     newProjectStatus(db, opts...),
		Repository:        newRepository(db, opts...),
		RepositoryAccount: newRepositoryAccount(db, opts...),
		Student:           newStudent(db, opts...),
		Task:              newTask(db, opts...),
		TaskStatus:        newTaskStatus(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	CalendarAccount   calendarAccount
	Meeting           meeting
	MeetingStatus     meetingStatus
	Professor         professor
	Project           project
	ProjectStage      projectStage
	ProjectStatus     projectStatus
	Repository        repository
	RepositoryAccount repositoryAccount
	Student           student
	Task              task
	TaskStatus        taskStatus
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		CalendarAccount:   q.CalendarAccount.clone(db),
		Meeting:           q.Meeting.clone(db),
		MeetingStatus:     q.MeetingStatus.clone(db),
		Professor:         q.Professor.clone(db),
		Project:           q.Project.clone(db),
		ProjectStage:      q.ProjectStage.clone(db),
		ProjectStatus:     q.ProjectStatus.clone(db),
		Repository:        q.Repository.clone(db),
		RepositoryAccount: q.RepositoryAccount.clone(db),
		Student:           q.Student.clone(db),
		Task:              q.Task.clone(db),
		TaskStatus:        q.TaskStatus.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                db,
		CalendarAccount:   q.CalendarAccount.replaceDB(db),
		Meeting:           q.Meeting.replaceDB(db),
		MeetingStatus:     q.MeetingStatus.replaceDB(db),
		Professor:         q.Professor.replaceDB(db),
		Project:           q.Project.replaceDB(db),
		ProjectStage:      q.ProjectStage.replaceDB(db),
		ProjectStatus:     q.ProjectStatus.replaceDB(db),
		Repository:        q.Repository.replaceDB(db),
		RepositoryAccount: q.RepositoryAccount.replaceDB(db),
		Student:           q.Student.replaceDB(db),
		Task:              q.Task.replaceDB(db),
		TaskStatus:        q.TaskStatus.replaceDB(db),
	}
}

type queryCtx struct {
	CalendarAccount   ICalendarAccountDo
	Meeting           IMeetingDo
	MeetingStatus     IMeetingStatusDo
	Professor         IProfessorDo
	Project           IProjectDo
	ProjectStage      IProjectStageDo
	ProjectStatus     IProjectStatusDo
	Repository        IRepositoryDo
	RepositoryAccount IRepositoryAccountDo
	Student           IStudentDo
	Task              ITaskDo
	TaskStatus        ITaskStatusDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		CalendarAccount:   q.CalendarAccount.WithContext(ctx),
		Meeting:           q.Meeting.WithContext(ctx),
		MeetingStatus:     q.MeetingStatus.WithContext(ctx),
		Professor:         q.Professor.WithContext(ctx),
		Project:           q.Project.WithContext(ctx),
		ProjectStage:      q.ProjectStage.WithContext(ctx),
		ProjectStatus:     q.ProjectStatus.WithContext(ctx),
		Repository:        q.Repository.WithContext(ctx),
		RepositoryAccount: q.RepositoryAccount.WithContext(ctx),
		Student:           q.Student.WithContext(ctx),
		Task:              q.Task.WithContext(ctx),
		TaskStatus:        q.TaskStatus.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}