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

	auth := router.Group("/auth") // localhost:8080/auth/sign-up
	{
		auth.POST("/sign-up", h.userIdentity, h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	return router
}
