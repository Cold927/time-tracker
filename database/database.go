package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time-tracker/config"
)

var Database *gorm.DB

func Connect(cfg config.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=Europe/Moscow",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password,
		cfg.Postgres.DbName, cfg.Postgres.SslMode)
	Database, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		Database.Exec("CREATE DATABASE IF NOT EXISTS timetracker")
	} else {
		fmt.Println("Успешное подключение к базе данных")
	}

	pgDb, _ := Database.DB()
	err = pgDb.Ping()
	if err != nil {
		panic(err)
	}

	pgDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	pgDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	pgDb.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

}

func GetDb() *gorm.DB {
	return Database
}

func CloseDb() {
	conn, _ := Database.DB()
	conn.Close()
}
