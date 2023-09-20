package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"tinkapi/v2/pkg/repository/entity"
)

func (h *Handler) createFromDesc(c *gin.Context) {
	var desc entity.Desc
	if err := c.ShouldBindYAML(&desc); err != nil {
		h.respondWithError(http.StatusBadRequest, c, err.Error())
		return
	}

	for _, team := range desc.Teams {
		dbTeam := entity.Team{
			Name:               team.Name,
			SchedulingTimezone: team.SchedulingTimezone,
			Email:              team.Email,
			SlackChannel:       team.SlackChannel,
		}

		if _, err := h.services.Team.CreateTeam(dbTeam); err != nil {
			h.respondWithError(http.StatusInternalServerError, c, err.Error())
			return
		}
		for _, user := range team.Users {
			dbUser := entity.User{
				Name:        user.Name,
				FullName:    user.FullName,
				PhoneNumber: user.PhoneNumber,
				Email:       user.Email,
			}

			if _, err := h.services.Authorization.CreateUser(dbUser); err != nil {
				h.respondWithError(http.StatusInternalServerError, c, err.Error())
				return
			}

			for _, duty := range user.Duty {
				dbDate, _ := time.Parse("02/01/2006", duty.Date)
				dbDuty := entity.Duty{
					User: dbUser,
					Team: dbTeam,
					Date: dbDate,
					Role: duty.Role,
				}
				if _, err := h.services.Duty.CreateDuty(dbDuty); err != nil {
					h.respondWithError(http.StatusInternalServerError, c, err.Error())
					return
				}
			}
		}
	}
	h.respondWithYAML(http.StatusOK, c, map[string]string{
		"status": "ok",
	})
}
