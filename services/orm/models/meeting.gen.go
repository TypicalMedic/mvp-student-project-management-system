// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameMeeting = "meeting"

// Meeting mapped from table <meeting>
type Meeting struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OrganizerID   int32     `gorm:"column:organizer_id;not null" json:"organizer_id"`
	ParticipantID int32     `gorm:"column:participant_id;not null" json:"participant_id"`
	Time          time.Time `gorm:"column:time;not null" json:"time"`
	IsOnline      bool      `gorm:"column:is_online;not null" json:"is_online"`
	StatusID      int32     `gorm:"column:status_id;not null" json:"status_id"`
	Description   string    `gorm:"column:description" json:"description"`
}

// TableName Meeting's table name
func (*Meeting) TableName() string {
	return TableNameMeeting
}
