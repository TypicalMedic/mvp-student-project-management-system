package managingstudentproject

import (
	"fmt"
	"time"
)

type MeetingStatus int

const (
	MeetingPlanned TaskStatus = iota
	MeetingPassed
	MeetingCancelled
)

func (s MeetingStatus) String() string {
	switch s {
	case MeetingStatus(MeetingPlanned):
		return "Planned"
	case MeetingStatus(MeetingPassed):
		return "Passed"
	case MeetingStatus(MeetingCancelled):
		return "Cancelled"
	default:
		return fmt.Sprintf("%d", int(s))
	}
}

type Meeting struct {
	id          uint
	organizer   Professor
	participant Student
	time        time.Time
	isOnline    bool
	status      MeetingStatus
	description string
}

func (m *Meeting) Id() uint {
	return m.id
}

func (m *Meeting) Organizer() Professor {
	return m.organizer
}

func (m *Meeting) Participant() Student {
	return m.participant
}

func (m *Meeting) Time() time.Time {
	return m.time
}

func (m *Meeting) IsOnline() bool {
	return m.isOnline
}

func (m *Meeting) Status() MeetingStatus {
	return m.status
}

func (m *Meeting) Description() string {
	return m.description
}

func (m *Meeting) SetId(id uint) {
	m.id = id
}

func (m *Meeting) SetOrganizer(organizer Professor) {
	m.organizer = organizer
}

func (m *Meeting) SetParticipant(participant Student) {
	m.participant = participant
}

func (m *Meeting) SetTime(time time.Time) {
	m.time = time
}

func (m *Meeting) SetIsOnline(isOnline bool) {
	m.isOnline = isOnline
}

func (m *Meeting) SetStatus(status MeetingStatus) {
	m.status = status
}

func (m *Meeting) SetDescription(description string) {
	m.description = description
}

func InitMeeting(id uint, participant Student, time time.Time, isOnline bool, status MeetingStatus, org Professor, desc string) Meeting {
	return Meeting{id: id, participant: participant, time: time, isOnline: isOnline, status: status, organizer: org, description: desc}
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
