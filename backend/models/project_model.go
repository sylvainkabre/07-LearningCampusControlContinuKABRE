package models

import (
	"time"

	"gorm.io/datatypes"
)

type Project struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Image       string
	Skills      datatypes.JSONSlice[string] `gorm:"type:json"`
}

type ProjectUpdateInput struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Image       *string   `json:"image"`
	Skills      *[]string `json:"skills"`
}
