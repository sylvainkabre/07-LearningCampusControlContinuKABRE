package models

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	ProjectID uint `json:"project_id"`
	UserID    uint
	Content   string
	CreateAt  string
	UpdatedAt string
}
