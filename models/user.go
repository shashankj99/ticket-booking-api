package models

import "time"

type UserRole string

const (
	Manager  UserRole = "manager"
	Attendee UserRole = "attendee"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-"`
	Role      UserRole  `json:"role" gorm:"type:enum('manager', 'attendee');default:'attendee'"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
