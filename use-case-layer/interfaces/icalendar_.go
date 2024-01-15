package interfaces

import (
	entities "mvp-student-project-management-system/managing-student-project"
)

type ICalendar interface {
	AddMeeting(entities.Meeting)
	RemoveMeeting(entities.Meeting)
	GetProfessorSchedule(entities.Professor) []entities.Meeting
}
