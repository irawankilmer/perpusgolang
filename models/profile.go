package models

import "time"

type Profile struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	DateOfBirth string    `json:"date_of_birth" gorm:"type:date"`
	Verified    string    `json:"verified" gorm:"type:enum('verified', 'waiting', 'blocked')"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
