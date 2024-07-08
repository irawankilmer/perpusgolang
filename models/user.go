package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique; required"`
	Email     string    `json:"email" gorm:"unique; required"`
	Password  string    `json:"password" gorm:"required"`
	Role      string    `json:"role" gorm:"type:enum('admin', 'librarian', 'member'); required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Profile   Profile   `json:"-"`
}
