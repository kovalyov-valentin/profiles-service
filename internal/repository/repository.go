package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/kovalyov-valentin/profiles-service/internal/models"
)

type UserServicer interface {
	CreateUser(ctx context.Context, user models.Users) (int, error)
	GetUser(ctx context.Context, id int) (models.Users, error)
	GetUsers(ctx context.Context) ([]models.Users, error)
	UpdateUser(ctx context.Context, id int, user models.Users) error
	DeleteUser(ctx context.Context, id int) error
}

type Repository struct {
	UserServicer
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserServicer: NewUserPostgres(db),
	}
}
