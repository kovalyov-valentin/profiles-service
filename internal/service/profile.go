package service

import "github.com/kovalyov-valentin/profiles-service/internal/repository"

type ProfileService struct {
	repo repository.ProfileServicer
}

func NewProfileService(repo repository.ProfileServicer) *ProfileService {
	return &ProfileService{
		repo: repo,
	}
}
