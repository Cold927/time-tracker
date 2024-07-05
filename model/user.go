package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"time-tracker/database"
	"time-tracker/utils"
)

type User struct {
	ID             uuid.UUID      `gorm:"primarykey" json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	Surname        string         `gorm:"size:255;not null;type:text" json:"surname" example:"Иванов"`
	Name           string         `gorm:"size:255;not null;type:text" json:"name" example:"Иван"`
	Patronymic     string         `gorm:"size:255;not null;type:text" json:"patronymic" example:"Иванович"`
	Address        string         `gorm:"size:255;not null;type:text" json:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
	PassportSeries int            `gorm:"uniqueIndex:idx_passport" json:"passportSeries"`
	PassportNumber int            `gorm:"uniqueIndex:idx_passport" json:"passportNumber"`
	Tasks          []Task         `json:"-"`
}

type UserCreate struct {
	Surname        string `json:"surname" example:"Иванов"`
	Name           string `json:"name" example:"Иван"`
	Patronymic     string `json:"patronymic" example:"Иванович"`
	Address        string `json:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
	PassportNumber string `json:"passportNumber" example:"1234 567890"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}

func (user User) Save(userCreate *UserCreate) (User, error) {
	passportSeries, passportNumber, err := utils.ParsePassport(userCreate.PassportNumber)
	if err != nil {
		return User{}, err
	}
	newUser := User{
		Surname:        userCreate.Surname,
		Name:           userCreate.Name,
		Patronymic:     userCreate.Patronymic,
		Address:        userCreate.Address,
		PassportNumber: passportNumber,
		PassportSeries: passportSeries,
	}
	if err := database.Database.Create(&newUser).Error; err != nil {
		return User{}, err
	}
	return newUser, nil
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
	if err := database.Database.Find(&user, "ID=?", uid).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
