package service

import (
	"context"
	"fmt"
	"github.com/kovalyov-valentin/profiles-service/internal/enrich"
	"github.com/kovalyov-valentin/profiles-service/internal/models"
	"github.com/kovalyov-valentin/profiles-service/internal/repository"
	"github.com/sirupsen/logrus"
)

type UserService struct {
	repo repository.UserServicer
	enrich.Enrichment
}

func NewUserService(repo repository.UserServicer, enrichment enrich.Enrichment) *UserService {
	return &UserService{
		Enrichment: enrichment,
		repo:       repo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user models.Users) (int, error) {
	age, err := s.Enrichment.GetAgeByName(user.Name)
	if err != nil {
		logrus.WithError(err).Error("Failed to enrich age")
		return 0, fmt.Errorf("failed to enrich age: %w", err)
	}
	user.Age = age

	gender, err := s.Enrichment.GetGenderByName(user.Name)
	if err != nil {
		logrus.WithError(err).Error("Failed to enrich gender")
		return 0, fmt.Errorf("failed to enrich gender: %w", err)
	}
	user.Gender = gender

	nationality, err := s.Enrichment.GetNationalityByName(user.Name)
	if err != nil {
		logrus.WithError(err).Error("Failed to enrich nationality")
		return 0, fmt.Errorf("failed to enrich nationality: %w", err)
	}
	user.Nationality = nationality

	userId, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		logrus.WithError(err).Error("Failed to create user")
		return 0, fmt.Errorf("failed to create user: %w", err)
	}
	user.Id = userId

	logrus.Debugf("Enriched and created user: %+v", user)

	return userId, nil

}

func (s *UserService) GetUser(ctx context.Context, id int) (models.Users, error) {
	return s.repo.GetUser(ctx, id)
}

func (s *UserService) GetUsers(ctx context.Context) ([]models.Users, error) {
	return s.repo.GetUsers(ctx)
}

func (s *UserService) UpdateUser(ctx context.Context, id int, user models.Users) error {
	return s.repo.UpdateUser(ctx, id, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.DeleteUser(ctx, id)
}
