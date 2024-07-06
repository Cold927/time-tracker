package migration

import (
	"gorm.io/gorm"
	"log"
	"time-tracker/model"
)

func StartTaskMigration(db *gorm.DB) {
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
