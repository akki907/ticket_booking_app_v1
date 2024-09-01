package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRole is an enum type used to restrict the values for user role.
type UserRole string

const (
	Manager  UserRole = "manager"
	attendee UserRole = "attendee"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email     string    `json:"email" gorm:"text;not null"`
	Role      UserRole  `json:"role" gorm:"text;default:attendee"`
	Password  string    `json:"-"` // Do not compute the password in json
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
	if u.Email == "ak.akki907@gmail.com" {
		db.Model(u).Update("role", Manager)
	}
	return
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}
