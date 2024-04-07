package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DatabaseConfig struct {
	DbHost string `envconfig:"HOST"`
	DbPort string `envconfig:"PORT"`
	DbName string `envconfig:"NAME"`
	DbUser string `envconfig:"USER"`
	DbPass string `envconfig:"PASS"`
}

func CreateConnection(config DatabaseConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DbHost, config.DbUser, config.DbPass, config.DbName, config.DbPort)
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
