package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	DBHost string `envconfig:"HOST"`
	DBPort string `envconfig:"PORT"`
	DBName string `envconfig:"NAME"`
	DBUser string `envconfig:"USER"`
	DBPass string `envconfig:"PASS"`
}

func CreateConnection(config *DatabaseConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost, config.DBUser, config.DBPass, config.DBName, config.DBPort)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
