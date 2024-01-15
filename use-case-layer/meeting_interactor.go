package usecaselayer

import (
	entities "mvp-student-project-management-system/managing-student-project"
	"mvp-student-project-management-system/use-case-layer/interfaces"
	irepositories "mvp-student-project-management-system/use-case-layer/interfaces/irepositories"
)

type MeetingtInteractor struct {
	MeetingDataStorage irepositories.IMeetingDataStorage
	Presenter          interfaces.IPresenter
	Calendar           interfaces.ICalendar
}

func (m *MeetingtInteractor) ArrangeMeeting(meeting entities.Meeting) {
	m.MeetingDataStorage.SaveMeeting(meeting)
	m.Calendar.AddMeeting(meeting)
}

func (m *MeetingtInteractor) CancelMeeting(meeting entities.Meeting) {
	meeting.SetStatusToCancelled()
	m.MeetingDataStorage.UpdateMeeting(meeting)
	m.Calendar.RemoveMeeting(meeting)
}

func (m *MeetingtInteractor) ViewProfSchedule(professor entities.Professor) {
	schedule := m.Calendar.GetProfessorSchedule(professor)
	m.Presenter.ShowSchedule(schedule)
}

func (m *MeetingtInteractor) DeleteMeeting(meeting entities.Meeting) {
	m.MeetingDataStorage.DeleteMeeting(meeting)
	m.Calendar.RemoveMeeting(meeting)
}
