package models

import "time"

type User struct {
	ID uint `gorm:"primaryKey"`

	Username string `gorm:"size:100;uniqueIndex;not null"`
	Password string `gorm:"not null"`

	MFAEnabled bool   `gorm:"default:false"`
	TOTPSecret string

	FailedAttempts int        `gorm:"default:0"`
	LockedUntil    *time.Time `gorm:"default:null"`

	LastLogin *time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}