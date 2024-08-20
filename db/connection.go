package db

import (
	"database/sql"
	"fmt"

	"github.com/HaikalRFadhilahh/auth/helper"
)

type DBConfig struct {
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
}

func EnvironmentDB() *DBConfig {
	return &DBConfig{
		DB_HOST:     helper.GetEnv("DB_HOST", "127.0.0.1"),
		DB_PORT:     helper.GetEnv("DB_PORT", "3306"),
		DB_USERNAME: helper.GetEnv("DB_USERNAME", "root"),
		DB_PASSWORD: helper.GetEnv("DB_PASSWORD", ""),
		DB_NAME:     helper.GetEnv("DB_NAME", "golang"),
	}
}

func InitDB(e *DBConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", e.DB_USERNAME, e.DB_PASSWORD, e.DB_HOST, e.DB_PORT, e.DB_NAME)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return db, err
}
