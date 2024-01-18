package meeting

import (
	entities "mvp-student-project-management-system/managing-student-project"
)

func SetOrganizer(org entities.Professor) func(*entities.Meeting) {
	return func(m *entities.Meeting) {
		m.SetOrganizer(org)
	}
}

func SetParticipant(participant entities.Student) func(*entities.Meeting) {
	return func(m *entities.Meeting) {
		m.SetParticipant(participant)
	}
}
