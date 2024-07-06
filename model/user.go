package model

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"time-tracker/database"
	"time-tracker/utils"
)

type User struct {
	ID             uuid.UUID      `gorm:"primarykey;default:uuid_generate_v4()" json:"id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	Surname        string         `gorm:"size:255;not null;type:text" json:"surname" example:"Иванов"`
	Name           string         `gorm:"size:255;not null;type:text" json:"name" example:"Иван"`
	Patronymic     string         `gorm:"size:255;not null;type:text" json:"patronymic" example:"Иванович"`
	Address        string         `gorm:"size:255;not null;type:text" json:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
	PassportSeries int            `gorm:"check:checker_passport_series,passport_series >= 1000 AND passport_series <= 9999;uniqueIndex:idx_passport;not null" json:"passportSeries"`
	PassportNumber int            `gorm:"check:checker_passport_number,passport_number >= 100000 AND passport_number <= 999999;uniqueIndex:idx_passport;not null" json:"passportNumber"`
	Tasks          []Task         `json:"-" gorm:"constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}

type UserCreate struct {
	Surname        string `json:"surname" example:"Иванов"`
	Name           string `json:"name" example:"Иван"`
	Patronymic     string `json:"patronymic" example:"Иванович"`
	Address        string `json:"address" example:"г. Москва, ул. Ленина, д. 5, кв. 1"`
	PassportNumber string `json:"passportNumber" example:"1234 567890"`
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

func (user User) UpdateData(uid string, data *UserCreate) (User, error) {
	passportSeries, passportNumber, err := utils.ParsePassport(data.PassportNumber)
	if err != nil {
		return User{}, err
	}
	newUser := User{
		Surname:        data.Surname,
		Name:           data.Name,
		Patronymic:     data.Patronymic,
		Address:        data.Address,
		PassportNumber: passportNumber,
		PassportSeries: passportSeries,
	}
	if err := database.Database.Where("ID=?", uid).Updates(&newUser).Error; err != nil {
		return User{}, err
	}
	return newUser, nil
}

func (user User) ListUsers(pagination utils.Pagination, field string, search string) (*utils.Pagination, error) {
	var users []*User
	if search != "" {
		database.Database.Scopes(utils.Paginate(&user, &pagination, database.Database)).Where(fmt.Sprintf(`%s ILIKE '%%%s%%'`, field, search)).Find(&users)
	} else {
		database.Database.Scopes(utils.Paginate(&user, &pagination, database.Database)).Find(&users)
	}
	pagination.Rows = users
	return &pagination, nil
}

func (user User) Info(series int, number int) (User, error) {
	if err := database.Database.Find(&user, "passport_series=? AND passport_number=?", series, number).Error; err != nil {
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
