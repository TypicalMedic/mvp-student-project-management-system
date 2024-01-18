package interfaces

import (
	entities "mvp-student-project-management-system/managing-student-project"
)

type IMeetingRepository interface {
	SaveMeeting(entities.Meeting)
	DeleteMeeting(entities.Meeting)
	UpdateMeeting(entities.Meeting)
}
