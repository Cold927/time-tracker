package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `envconfig:"APP_PORT"`
	Cors struct {
		AllowOrigins string `envconfig:"CORS"`
	}
	Postgres struct {
		Host         string `envconfig:"DB_HOST"`
		Port         int    `envconfig:"DB_PORT"`
		User         string `envconfig:"DB_USER"`
		Password     string `envconfig:"DB_PASSWORD"`
		DbName       string `envconfig:"DB_NAME"`
		SslMode      string `envconfig:"DB_SSL_MODE"`
		MaxIdleConns int    `envconfig:"DB_MAX_IDLE_CONNECTIONS"`
		MaxOpenConns int    `envconfig:"DB_MAX_OPEN_CONNECT"`
	}
}

var (
	AppConfig Config // Экспортируем переменную AppConfig для доступа из других пакетов
)

func init() {
	GetConfig()
}

func GetConfig() (Config, error) {
	_ = godotenv.Load(".env")
	cnf := Config{}
	err := envconfig.Process("", &cnf)
	return cnf, err
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "/app/"
	} else if env == "production" {
		return ".env.production"
	} else {
		return "$HOME/.env"
	}
}
