package managingstudentproject

import (
	"time"
)

type MeetingStatus int

const (
	MeetingPlanned TaskStatus = iota
	MeetingPassed
	MeetingCancelled
)

type Meeting struct {
	student  Student
	time     time.Time
	isOnline bool
	status   MeetingStatus
}

func InitMeeting(student Student, time time.Time, isOnline bool, status MeetingStatus) Meeting {
	return Meeting{student: student, time: time, isOnline: isOnline, status: status}
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
