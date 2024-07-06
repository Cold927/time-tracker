package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"time-tracker/database"
)

const (
	ACTIVE    string = "ACTIVE"
	COMPLETED string = "COMPLETED"
)

type Task struct {
	ID          uuid.UUID      `gorm:"primarykey;default:uuid_generate_v4()" json:"id"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Title       string         `gorm:"type:text" json:"title" example:"Новая задача"`
	Description string         `gorm:"type:text" json:"description" example:"Описание задачи"`
	StartDate   time.Time      `gorm:"not null;" json:"start_date"`
	EndDate     time.Time      `gorm:"not null;" json:"end_date"`
	TotalTime   int            `gorm:"not null;" json:"total_time"`
	Status      string         `gorm:"not null;" json:"status"`
	UserID      uuid.UUID      `gorm:"not null;" json:"-"`
}

type TaskCreate struct {
	Title       string `gorm:"type:text" json:"title" example:"Новая задача"`
	Description string `gorm:"type:text" json:"description" example:"Описание задачи"`
}

type TaskResponse struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	PeriodTime  string    `json:"period_time"`
	Status      string    `json:"status"`
}

func (task Task) Info(uid string, startDate string, endDate string) ([]*TaskResponse, error) {
	var tasks []*TaskResponse
	err := database.Database.Model(&task).
		Select(`tasks.title,
			tasks.description,
			tasks.start_date,
			tasks.end_date,
			CONCAT(
	       FLOOR(SUM(EXTRACT(EPOCH FROM (end_date - start_date))) / (24 * 3600)), ':',
	       CASE WHEN FLOOR((SUM(EXTRACT(EPOCH FROM (end_date - start_date))) % (24 * 3600)) / 3600) < 10 THEN '0' ELSE '' END,
	       FLOOR((SUM(EXTRACT(EPOCH FROM (end_date - start_date))) % (24 * 3600)) / 3600), ':',
	       CASE WHEN FLOOR((SUM(EXTRACT(EPOCH FROM (end_date - start_date))) % 3600) / 60) < 10 THEN '0' ELSE '' END,
	       FLOOR((SUM(EXTRACT(EPOCH FROM (end_date - start_date))) % 3600) / 60), ':',
	       CASE WHEN (SUM(EXTRACT(EPOCH FROM (end_date - start_date))) % 60) < 10 THEN '0' ELSE '' END,
	       (ROUND(SUM(EXTRACT(EPOCH FROM (end_date - start_date))) % 60))
	   ) AS period_time, tasks.status`).
		Where("user_id = ?", uid).
		Where("start_date >= ?", startDate).
		Where("end_date <= ?", endDate).
		Where("status = 'COMPLETED'").
		Group("id, title, description, start_date, end_date").
		Order("total_time DESC").
		Find(&tasks).Error
	if err != nil {
		return []*TaskResponse{}, err
	}
	return tasks, nil
}

func (task Task) CountdownStart() (Task, error) {
	task.StartDate = time.Now()
	task.EndDate = time.Now()
	task.Status = ACTIVE
	if err := database.Database.Create(&task).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (task Task) CountdownEnd(tid string) (Task, error) {
	task.Status = COMPLETED
	task.EndDate = time.Now()
	err := database.Database.Where("ID=?", tid).Updates(&task).Error
	if err != nil {
		return Task{}, err
	}
	return task, nil
}
