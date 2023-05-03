package database

import (
	"time"

	"gorm.io/gorm"
)

// TODO: Set gorm db types

// swagger:model User
type User struct {
	ID        uint            `gorm:"primary_key;<-:create" json:"id"`
	CreatedAt time.Time       `gorm:"<-:create" json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Password  string          `json:"password" gorm:"-"` // Ignore
}

// swagger:model Schedule
type Schedule struct {
	ID            uint            `gorm:"primary_key;<-:create" json:"id"`
	CreatedAt     time.Time       `gorm:"<-:create" json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	DeletedAt     *gorm.DeletedAt `json:"deleted_at"`
	Title         string          `json:"title"`
	Description   string          `json:"description"`
	StartDateTime string          `json:"start_date_time"`
	EndDateTime   string          `json:"end_date_time"`
	Speakers      []Speaker       `json:"speakers"`
}

// swagger:model Speaker
type Speaker struct {
	ID          uint      `gorm:"primary_key;<-:create" json:"id"`
	CreatedAt   time.Time `gorm:"<-:create" json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Topic       string    `json:"topic"`
	Description string    `json:"description"`
	Audience    string    `json:"audience"`
	Order       int       `json:"order"`
	Minutes     int       `json:"minutes"`
	ScheduleId  int       `json:"schedule_id"`
	Private     bool      `gorm:"not null" json:"private"`
}

// swagger:model Comment
type Comment struct {
	ID          uint            `gorm:"primary_key;<-:create" json:"id"`
	CreatedAt   time.Time       `gorm:"<-:create" json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `json:"deleted_at"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
}
