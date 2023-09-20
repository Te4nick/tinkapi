package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	team := router.Group("/team")
	{
		team.POST("/create-from-desc", h.createFromDesc)
	}

	return router
}

func (h *Handler) respondWithError(code int, c *gin.Context, msg string) {
	logrus.Error(msg)
	c.AbortWithStatusJSON(code, msg)
}

func (h *Handler) respondWithYAML(code int, c *gin.Context, payload interface{}) {
	c.YAML(code, payload)
}
