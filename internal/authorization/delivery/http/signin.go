package v1

import (
	"context"
	"net/http"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

// Signin godoc
// @Description  signin service with username and password  set Header : access_token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        input    body      models.SigninCreds  true  "user creds"
// @Success      200      {object}  models.Response
// @Failure      400,500  {object}  models.Response
// @Router       /v1/auth/signin [post]
func (h *Handler) SignIn(c *gin.Context) {
	var (
		user  *models.SigninCreds
		err   error
		res   models.User
		token *models.TokenDetails
	)

	err = c.ShouldBindJSON(&user)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	res, err = h.authUseCases.SignIn(ctx, user.Username, user.Password)
	if err != nil {
		h.logger.Error(err)
		if err.Error() == "sql: no rows in result set" {
			responseWithStatus(c, http.StatusBadRequest, err.Error(), "Incorrect password or username", nil)
			return
		}
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "Internal server error", nil)
		return
	}

	// call signin/signup/refresh
	token, err = CreateToken(res.ID, h.cfg.App.SecretAccess, h.cfg.App.SecretRefresh)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	token.UserID = res.ID

	// err = h.useCases.CreateSession(ctx, token)
	// refresh, /signin call
	// err = h.useCases.UpdateSession(ctx, token)

	c.Writer.Header().Set("refresh_token", token.RefreshToken)
	c.Writer.Header().Set("access_token", token.AccessToken)
	responseWithStatus(c, http.StatusOK, "success signin", "OK", res)
}
