// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProfessor = "professor"

// Professor mapped from table <professor>
type Professor struct {
	ID                int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name              string `gorm:"column:name;not null" json:"name"`
	Surname           string `gorm:"column:surname;not null" json:"surname"`
	MiddleName        string `gorm:"column:middle_name;not null" json:"middle_name"`
	Position          string `gorm:"column:position;not null" json:"position"`
	CalendarAccountID int32  `gorm:"column:calendar_account_id" json:"calendar_account_id"`
	CalendarID        string `gorm:"column:calendar_id" json:"calendar_id"`
	RepoAccountID     int32  `gorm:"column:repo_account_id" json:"repo_account_id"`
	CloudDriveFolderID     string  `gorm:"column:cloud_drive_folder_id" json:"cloud_drive_folder_id"`
}

// TableName Professor's table name
func (*Professor) TableName() string {
	return TableNameProfessor
}
