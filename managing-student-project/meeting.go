package managingstudentproject

import (
	"time"

	"github.com/google/uuid"
)

type MeetingStatus int

const (
	MeetingPlanned TaskStatus = iota
	MeetingPassed
	MeetingCancelled
)

type Meeting struct {
	id          uuid.UUID
	organizer   Professor
	participant Student
	time        time.Time
	isOnline    bool
	status      MeetingStatus
}

func InitMeeting(participant Student, time time.Time, isOnline bool, status MeetingStatus, org Professor) Meeting {
	return Meeting{id: uuid.New(), participant: participant, time: time, isOnline: isOnline, status: status, organizer: org}
}

// add lower interface
func (m *Meeting) SetStatusToPlanned() {
	m.status = MeetingStatus(MeetingPlanned)
}

// add lower interface
func (m *Meeting) SetStatusToPassed() {
	m.status = MeetingStatus(MeetingPassed)
}

// add lower interface
func (m *Meeting) SetStatusToCancelled() {
	m.status = MeetingStatus(MeetingCancelled)
}
