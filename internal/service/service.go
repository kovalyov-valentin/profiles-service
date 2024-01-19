package service

import (
	"context"
	"github.com/kovalyov-valentin/profiles-service/internal/config"
	"github.com/kovalyov-valentin/profiles-service/internal/enrich"
	"github.com/kovalyov-valentin/profiles-service/internal/models"
	"github.com/kovalyov-valentin/profiles-service/internal/repository"
)

type UserServicer interface {
	CreateUser(ctx context.Context, user models.Users) (int, error)
	GetUser(ctx context.Context, id int) (models.Users, error)
	GetUsers(ctx context.Context) ([]models.Users, error)
	UpdateUser(ctx context.Context, id int, user models.Users) error
	DeleteUser(ctx context.Context, id int) error
}

type Service struct {
	UserServicer
}

func NewService(repos *repository.Repository, conf *config.Config) *Service {
	enrichment := enrich.NewEnrichment(conf)
	return &Service{
		UserServicer: NewUserService(repos.UserServicer, enrichment),
	}
}
