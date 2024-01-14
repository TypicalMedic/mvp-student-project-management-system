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
	id          int
	organizer   Professor
	participant Student
	time        time.Time
	isOnline    bool
	status      MeetingStatus
	description string
}

func (m *Meeting) GetId() int {
	return m.id
}

func (m *Meeting) GetOrganizer() Professor {
	return m.organizer
}

func (m *Meeting) GetParticipant() Student {
	return m.participant
}

func (m *Meeting) GetTime() time.Time {
	return m.time
}

func (m *Meeting) GetIsOnline() bool {
	return m.isOnline
}

func (m *Meeting) GetStatus() MeetingStatus {
	return m.status
}

func (m *Meeting) GetDescription() string {
	return m.description
}

func (m *Meeting) SetId(id int) {
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

func InitMeeting(participant Student, time time.Time, isOnline bool, status MeetingStatus, org Professor, desc string) Meeting {
	return Meeting{participant: participant, time: time, isOnline: isOnline, status: status, organizer: org, description: desc}
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
