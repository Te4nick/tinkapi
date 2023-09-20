package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"tinkapi/v2/pkg/repository/entity"
)

func (h *Handler) createUser(c *gin.Context) {
	var user entity.User
	err := c.ShouldBindYAML(&user)
	if err != nil {
		h.respondWithError(http.StatusBadRequest, c, err.Error())
		return
	}
	userName, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		h.respondWithError(http.StatusInternalServerError, c, err.Error())
		return
	}
	h.respondWithYAML(http.StatusOK, c, map[string]interface{}{
		"name": userName,
	})
	logrus.Info(user)
}

func (h *Handler) removeUser(c *gin.Context) {

}
