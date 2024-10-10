package config

import (
	"os"
	_ "github.com/joho/godotenv/autoload"
)


type MySqlConfig struct {
	Username string
	Password string
	URL      string
	Port     string
	Database string
}

type Config struct {
	mysqlConfig MySqlConfig
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		mysqlConfig: MySqlConfig{		
			Username: os.Getenv("MYSQL_USERNAME"),
			Password: os.Getenv("MYSQL_PASSWORD"),
			URL: os.Getenv("MYSQL_URL"),
			Port: os.Getenv("MYSQL_PORT"),
			Database: os.Getenv("MYSQL_DB"),
		},
	}

	return cfg, nil
}