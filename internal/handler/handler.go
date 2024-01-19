package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kovalyov-valentin/profiles-service/internal/config"
	"github.com/kovalyov-valentin/profiles-service/internal/service"
)

type Handler struct {
	services *service.Service
	config   *config.Config
}

func NewHandler(services *service.Service, config *config.Config) *Handler {
	return &Handler{
		services: services,
		config:   config,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/createUser", h.CreateUser)
		api.GET("/getUser/:id", h.GetUser)
		api.GET("/getUsers", h.GetUsers)
		api.PUT("/updateUser/:id", h.UpdateUser)
		api.DELETE("/deleteUser/:id", h.DeleteUser)
	}
	return router
}
