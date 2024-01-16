package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	config "github.com/kovalyov-valentin/profiles-service/internal/config"
)

func NewPostgresDB(cfg config.DB) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
