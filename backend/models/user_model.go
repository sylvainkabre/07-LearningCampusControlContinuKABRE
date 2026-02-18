package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string `gorm:"unique" binding:"required,email"`
	Password  string `gorm:"binding:required,min=6"`
}
