package model

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
	"time-tracker/database"
)

type Task struct {
	ID          uuid.UUID      `gorm:"primarykey" json:"-"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Title       string         `gorm:"type:text" json:"title" example:"Новая задача"`
	Description string         `gorm:"type:text" json:"description" example:"Описание задачи"`
	StartDate   datatypes.Date `gorm:"not null;" json:"-"`
	EndDate     datatypes.Date `gorm:"not null;" json:"-"`
	TotalTime   datatypes.Time `gorm:"not null;" json:"-"`
	UserID      uuid.UUID      `gorm:"not null;" json:"-"`
}

func (task Task) BeforeCreate(tx *gorm.DB) (err error) {
	task.ID = uuid.New()
	return
}

func (task Task) CountdownStart() (Task, error) {
	task.StartDate = datatypes.Date(time.Now())
	if err := database.Database.Create(&task).Error; err != nil {
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
