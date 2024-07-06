package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time-tracker/config"
)

var Database *gorm.DB

func Connect(cfg config.Config) {
	var err error

	dsnInit := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=%s TimeZone=Europe/Moscow",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password,
		cfg.Postgres.SslMode)
	dsn := fmt.Sprintf("%s dbname=%s", dsnInit, cfg.Postgres.DbName)
	Database, err = gorm.Open(postgres.Open(dsnInit))
	count := 0
	Database.Raw("SELECT count(*) FROM pg_database WHERE datname = ?", cfg.Postgres.DbName).Scan(&count)
	if count == 0 {
		sql := fmt.Sprintf("CREATE DATABASE %s", cfg.Postgres.DbName)
		Database.Exec(sql)
	}
	Database, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatal("Ошибка при открытии базы данных: ", err)
	} else {
		log.Println("Успешное подключение к базе данных")
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
