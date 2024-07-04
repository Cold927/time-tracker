package model

import (
	"gorm.io/gorm"
	"time-tracker/database"
)

type User struct {
	gorm.Model `json:"-"`
	Surname    string `json:"surname" gorm:"size:255;not null;type:text" example:"Иванов"`
	Name       string `json:"name" gorm:"size:255;not null;type:text" example:"Иван"`
	Patronymic string `json:"patronymic" gorm:"size:255;not null;type:text" example:"Иванович"`
	Address    string `json:"address" gorm:"size:255;not null;type:text" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
	Tasks      []Task `json:"-" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}

func (user User) Save() (User, error) {
	if err := database.Database.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (user User) UpdateData(uid string, data User) (User, error) {
	if err := database.Database.Where("ID=?", uid).Updates(data).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
func (user User) DeleteUser(uid string) (User, error) {
	if err := database.Database.Delete(&Task{}, "User_ID=?", uid).Error; err != nil {
		return User{}, err
	}

	if err := database.Database.Delete(&user, "ID=?", uid).Error; err != nil {
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
