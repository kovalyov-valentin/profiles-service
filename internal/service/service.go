package service

import "github.com/kovalyov-valentin/profiles-service/internal/repository"

type ProfileServicer interface {
}

type Service struct {
	Repo ProfileServicer
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Repo: repos,
	}
}
