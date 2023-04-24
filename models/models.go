package models

import (
	// "database/sql"
	"gorm.io/gorm"
)

// TODO: Set gorm db types

// swagger:model User
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password" gorm:"-"` // Ignore
}

// swagger:model ScheduleStatus
type ScheduleStatus struct {
	gorm.Model
	Title string `json:"title"`
}

// swagger:model Schedule
type Schedule struct {
	gorm.Model
	Title            string `json:"title"`
	StartDateTime    string `json:"start_date_time"`
	EndDateTime      string `json:"end_date_time"`
	ScheduleStatusId int    `json:"schedule_status_id" gorm:"foreignKey:ScheduleStatusId"`
	ScheduleStatus   ScheduleStatus
}

// swagger:model SpeakerStatus
type SpeakerStatus struct {
	gorm.Model
	Title string `json:"title"`
}

// swagger:model Speaker
type Speaker struct {
	gorm.Model
	Name            string `json:"name"`
	Topic           string `json:"topic"`
	Description     string `json:"description"`
	SpeakerStatusId int    `json:"speaker_status_id" gorm:"foreignKey:SpeakerStatusId"`
	SpeakerStatus   SpeakerStatus
}

// swagger:model Comment
type Comment struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
