package handler

import (
	"github.com/gin-gonic/gin"
	"ims-authentication-api/pkg/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/auth")
	{
		entities := api.Group("/user")
		{
			entities.POST("/sign-up", h.signUp)
			entities.GET("/sign-in", h.signIn)
		}
	}

	return router
}
