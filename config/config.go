package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Config struct {
	RunMode string `envconfig:"GIN_MODE"`
	Port    string `envconfig:"APP_PORT"`
	Cors    struct {
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
	filePath := getConfigPath(os.Getenv("APP_ENV"))
	_ = godotenv.Load(filePath)
	cnf := Config{}
	err := envconfig.Process("", &cnf)
	return cnf, err
}

func getConfigPath(env string) string {
	if env == "docker" {
		return "/app/.env"
	} else if env == "production" {
		return ".env.production"
	} else {
		return ".env.local"
	}
}
