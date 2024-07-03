package model

import (
	"gorm.io/gorm"
	"time-tracker/database"
)

type User struct {
	gorm.Model `json:"-"`
	Surname    string `json:"surname" gorm:"size:255;not null;"`
	Name       string `json:"name" gorm:"size:255;not null;"`
	Patronymic string `json:"patronymic" gorm:"size:255;not null;"`
	Address    string `json:"address" gorm:"size:255;not null;"`
	Tasks      []Task `json:"-"`
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}
