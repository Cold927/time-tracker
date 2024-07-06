package migration

import (
	"gorm.io/gorm"
	"log"
	"time-tracker/model"
)

func StartTaskMigration(db *gorm.DB) {
	// Создание типа status_task в бд для статуса заданий
	var exists bool
	if err := db.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_task')").Scan(&exists).Error; err != nil {
		log.Fatal("status_task уже существует")
	}
	if !exists {
		if err := db.Exec("CREATE TYPE status_task AS ENUM ('ACTIVE', 'COMPLETED')").Error; err != nil {
			log.Fatal("status_task уже существует")
		}
	}
	// Создание таблицы задач
	if err := db.AutoMigrate(&model.Task{}); err != nil {
		log.Fatalf("Ошибка при миграции таблицы Task: %v", err)
	}

	// Функция расчета затрат времени на задачу
	if err := db.Exec(`
        CREATE OR REPLACE FUNCTION calculate_total_time()
        RETURNS TRIGGER AS $$
        BEGIN
            NEW.total_time := EXTRACT(EPOCH FROM (NEW.end_date - NEW.start_date));
            RETURN NEW;
        END;
        $$ LANGUAGE plpgsql;
    `).Error; err != nil {
		log.Fatal("Невозможно создать функцию:", err)
	}

	// Создание триггера calculate_total_time_trigger на таблице tasks
	if err := db.Exec(`
        CREATE OR REPLACE TRIGGER calculate_total_time_trigger
        BEFORE INSERT OR UPDATE ON tasks
        FOR EACH ROW
        WHEN (NEW.end_date IS NOT NULL)
        EXECUTE FUNCTION calculate_total_time();
    `).Error; err != nil {
		log.Fatal("Невозможно создать триггер:", err)
	}
}
