package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Surname    string `gorm:"size:255;not null;" json:"surname"`
	Name       string `gorm:"size:255;not null;" json:"name"`
	Patronymic string `gorm:"size:255;not null;" json:"patronymic"`
	Address    string `gorm:"size:255;not null;" json:"address"`
	Tasks      []Task
}
