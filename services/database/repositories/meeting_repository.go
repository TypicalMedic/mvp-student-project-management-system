package repositories

import (
	"log"
	entities "mvp-student-project-management-system/managing-student-project"
	"mvp-student-project-management-system/services/database"
	"mvp-student-project-management-system/services/database/datagateway/todatabase"
)

type DBMeetingRepository struct {
	DBContext database.Database
}

func (r *DBMeetingRepository) SaveMeeting(meeting entities.Meeting) {
	dbmeeting := todatabase.Meeting(meeting)
	err := r.DBContext.DBInteractor.Meeting.Create(&dbmeeting)
	if err != nil {
		log.Fatal(err)
	}
}
func (r *DBMeetingRepository) DeleteMeeting(meeting entities.Meeting) {
	dbmeeting := todatabase.Meeting(meeting)
	_, err := r.DBContext.DBInteractor.Meeting.Delete(&dbmeeting)
	if err != nil {
		log.Fatal(err)
	}
}
func (r *DBMeetingRepository) UpdateMeeting(meeting entities.Meeting) {
	dbmeeting := todatabase.Meeting(meeting)
	err := r.DBContext.DBInteractor.Meeting.Save(&dbmeeting)
	if err != nil {
		log.Fatal(err)
	}
}
