package v1

import (
	"context"
	"log"
	"net/http"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfileBio(c *gin.Context) {
	var (
		err    error
		result *models.Profile
	)
	userID, ok := c.Get("user_id")
	if !ok {
		log.Println("not exist userId value")
		responseWithStatus(c, http.StatusBadRequest, "refresh token incorrect", "no have userI_id, to refresh token", nil)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	result, err = h.profileUseCases.GetBioByUserID(ctx, int(userID.(float64)))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "Internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "succes return bio user", "OK", result)
}
