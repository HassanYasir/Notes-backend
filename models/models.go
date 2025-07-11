package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UserID      uint      `gorm:"not null" json:"user_id"` // Foreign key
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Tag         string    `gorm:"size:50;default:'General'" json:"tag"`
	Date        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"date"`
}

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:255;not null"`
	Notes    []Note // One-to-Many relationship
}

type RecievedNote struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
}
type SendingNotes struct {
	Notes []Note `json:"notes"`
}
