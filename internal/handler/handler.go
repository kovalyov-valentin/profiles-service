package handler

import (
	"github.com/kovalyov-valentin/profiles-service/internal/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}
func (h *Handler) InitRoutes() http.Handler {
	router := http.NewServeMux()

	return router
}
