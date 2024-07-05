package model

import (
	"database/sql/driver"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"time-tracker/database"
)

type statusTask string

const (
	ACTIVE    statusTask = "ACTIVE"
	COMPLETED statusTask = "COMPLETED"
)

func (status *statusTask) Scan(value interface{}) error {
	*status = statusTask(value.([]byte))
	return nil
}

func (status statusTask) Value() (driver.Value, error) {
	return string(status), nil
}

type Task struct {
	ID          uuid.UUID      `gorm:"primarykey;default:uuid_generate_v4()" json:"-"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Title       string         `gorm:"type:text" json:"title" example:"Новая задача"`
	Description string         `gorm:"type:text" json:"description" example:"Описание задачи"`
	StartDate   time.Time      `gorm:"not null;" json:"-"`
	EndDate     time.Time      `gorm:"not null;" json:"-"`
	TotalTime   int            `gorm:"not null;" json:"-"`
	Status      statusTask     `gorm:"type:status_task;not null;" json:"-"`
	UserID      uuid.UUID      `gorm:"not null;" json:"-"`
}

func (task *Task) BeforeCreate(tx *gorm.DB) (err error) {
	task.Status = ACTIVE
	return
}

func (task Task) CountdownStart() (Task, error) {
	task.StartDate = time.Now()
	task.EndDate = time.Now()
	if err := database.Database.Create(&task).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (task Task) CountdownEnd(tid string) (Task, error) {
	err := database.Database.Model(&task).Where("ID=?", tid).Updates(&Task{EndDate: time.Now(), Status: COMPLETED}).Error
	if err != nil {
		return Task{}, err
	}
	return task, nil
}
