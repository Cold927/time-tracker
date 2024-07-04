package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
	"time-tracker/database"
)

type Task struct {
	gorm.Model  `json:"-"`
	Description string         `gorm:"type:text" json:"description" example:"Новая задача"`
	StartDate   datatypes.Date `gorm:"not null;" json:"-"`
	EndDate     datatypes.Date `gorm:"not null;" json:"-"`
	TotalTime   datatypes.Time `gorm:"not null;" json:"-"`
	UserID      uint           `gorm:"not null;" json:"-"`
}

func (task Task) CountdownStart() (Task, error) {
	task.StartDate = datatypes.Date(time.Now())

	err := database.Database.Create(&task).Error
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (task Task) CountdownEnd(tid string) (Task, error) {
	err := database.Database.Model(&task).Where("ID=?", tid).Update("end_date", datatypes.Date(time.Now())).Error
	if err != nil {
		return Task{}, err
	}
	return task, nil
}
