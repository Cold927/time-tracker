package model

import (
	"gorm.io/gorm"
	"time-tracker/database"
)

type User struct {
	gorm.Model `json:"-"`
	Surname    string `json:"surname" gorm:"size:255;not null;" example:"Иванов"`
	Name       string `json:"name" gorm:"size:255;not null;" example:"Иван"`
	Patronymic string `json:"patronymic" gorm:"size:255;not null;" example:"Иванович"`
	Address    string `json:"address" gorm:"size:255;not null;" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
	Tasks      []Task `json:"-"`
}

func (user User) Save() (User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (user User) UpdateData(uid string, data any) (User, error) {
	err := database.Database.Where("ID=?", uid).Updates(data).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (user User) FindUserById(uid string) (User, error) {
	err := database.Database.Find(&user, "ID=?", uid).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
