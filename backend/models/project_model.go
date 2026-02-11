package models

import "time"

type Project struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	Image       string
	Skills      []string `gorm:"type:json"`
}
