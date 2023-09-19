package handler

import (
	"github.com/gin-gonic/gin"
	"tinkapi/v2/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/create-user", h.createUser)
		auth.POST("/remove-user", h.removeUser)
	}

	return router
}
