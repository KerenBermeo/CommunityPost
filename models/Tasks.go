package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string
	Description string
	Done        bool `gorm:"default:false"`
	ProjectID   uint
	UserID      uint
}
