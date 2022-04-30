package v1

import (
	"context"
	"log"
	"net/http"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

// or  sql query - getCreatedMachines ?

// Profile godoc
// @Description  Profile service, user info(bio)
// @Tags         Profile
// @Produce      json
// @Security BearerAuth
// @Failure      400,500  {object}  models.Response
// @Success      200      {object}  models.Profile
// @Router       /v1/profile [get]
func (h *Handler) GetProfileBio(c *gin.Context) {
	var (
		err    error
		result *models.Profile
	)
	userID, ok := c.Get("user_id")
	if !ok {
		log.Println("not exist userId value")
		responseWithStatus(c, http.StatusBadRequest, "access token incorrect", "no have userId, by token", nil)
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
