package v1

import (
	"context"
	"net/http"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

// todo: - token expires & user_id in redis -> check token_uuid;

// Signup godoc
// @Description  signup service with model User, set Header : access_token
// @Tags         Signup
// @Accept       json
// @Produce      json
// @Param        input    body      models.User  true  "user data"
// @Success      200      {object}  models.Response
// @Failure      400,500  {object}  models.Response
// @Router       /v1/auth/signup [post]
func (h *Handler) SignUp(c *gin.Context) {
	var (
		user   *models.User
		err    error
		lastID int
		token  *models.TokenDetails
	)

	err = c.ShouldBindJSON(&user)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lastID, err = h.authUseCases.SignUp(ctx, user)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}

	token, err = CreateToken(lastID, h.cfg.App.SecretAccess, h.cfg.App.SecretRefresh)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}

	c.Writer.Header().Set("refresh_token", token.RefreshToken)
	c.Writer.Header().Set("access_token", token.AccessToken)

	responseWithStatus(c, http.StatusOK, "success signup", "OK", lastID)
}
