package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// remove client & db jwt token; by user_id

// Logout godoc
// @Description  Logout service, recieve access_token, header remove - access, refresh token
// @Tags         Auth
// @Produce      json
// @Security BearerAuth
// @Success      200      {object}  models.Response
// @Router       /v1/auth/logout [post]
func (h *Handler) Logout(c *gin.Context) {
	c.Writer.Header().Set("refresh_token", "")
	c.Writer.Header().Set("access_token", "")
	responseWithStatus(c, http.StatusOK, "Success logout, remove tokens", "OK", nil)
}
