package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model `json:"-"`
	Content    string `gorm:"type:text" json:"content"`
	UserID     uint   `json:"-"`
}
