package v1

import (
	"net/http"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

// refresh - invalid -> redirect signin; else newToken

// RefreshJwt godoc
// @Description  refresh jwt token, recieve  Authorization : access_token, set Header new pair [access, refresh]
// @Tags         Auth
// @Produce      json
// @Security BearerAuth
// @Success      200      {object}  models.Response
// @Failure      400,500  {object}  models.Response
// @Router       /v1/auth/refresh [post]
func (h *Handler) RefreshJwt(c *gin.Context) {
	var (
		err   error
		token *models.TokenDetails
	)

	userID, ok := c.Get("user_id")
	if !ok {
		responseWithStatus(c, http.StatusBadRequest, "refresh token incorrect", "no have userI_id, to refresh token", nil)
		return
	}

	token, err = CreateToken(userID.(int64), h.cfg.App.SecretAccess, h.cfg.App.SecretRefresh)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	c.Writer.Header().Set("refresh_token", token.RefreshToken) // todo set cookie
	c.Writer.Header().Set("access_token", token.AccessToken)
	// models.Response // resp.Error, resp.Success

	responseWithStatus(c, http.StatusCreated, "refresh token created", "Created", nil)
}

// saveDbToken, delete, etc
