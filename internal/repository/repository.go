package repository

import "github.com/jmoiron/sqlx"

type ProfileServicer interface {
}

type Repository struct {
	ProfileServicer
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ProfileServicer: NewProfilePostgres(db),
	}
}
