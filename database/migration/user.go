package migration

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"time-tracker/model"
)

func StartUserMigration(db *gorm.DB) {
	// Создание таблицы пользователя
	db.AutoMigrate(&model.User{})
	// Создание проверки номера паспорта в бд
	if !db.Migrator().HasConstraint(&model.User{}, "checker_passport_number") {
		db.Migrator().CreateConstraint(&model.User{}, "checker_passport_number")
	}
	// Создание проверки серии паспорта в бд
	if !db.Migrator().HasConstraint(&model.User{}, "checker_passport_series") {
		db.Migrator().CreateConstraint(&model.User{}, "checker_passport_series")
	}
	if db.Find(&model.User{}).RowsAffected == 0 {
		filePath, err := filepath.Abs("database/migration/users.json")
		jsonFile, err := os.Open(filePath)
		if err != nil {
			log.Fatalf("Ошибка открытия файла: %v", err)
		}
		defer jsonFile.Close()
		jsonData, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Ошибка чтения файла: %v", err)
		}

		tx := db.Begin()
		if tx.Error != nil {
			log.Fatalf("Ошибка в начале транзакции")
		}

		var users []model.User
		if err := json.Unmarshal(jsonData, &users); err != nil {
			log.Fatalf("Ошибка форматирования данных: %v", err)
		}

		for _, user := range users {
			result := tx.Create(&user)
			if result.Error != nil {
				tx.Rollback() // Откатываем транзакцию при ошибке сохранения пользователя
				log.Printf("Ошибка при сохранении пользователя %s %s: %v", user.Name, user.Surname, result.Error)
				return // Прерываем выполнение, чтобы не продолжать сохранение других пользователей
			}
			fmt.Printf("Пользователь %s %s успешно сохранен\n", user.Name, user.Surname)
		}
		err = tx.Commit().Error
		if err != nil {
			log.Fatalf("Ошибка при фиксации транзакции: %v", err)
		}
	}
}
