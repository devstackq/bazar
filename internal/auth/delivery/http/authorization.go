package v1

import (
	"context"
	"net/http"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(c *gin.Context) {

	var (
		user *models.User
		err  error
		userID int
		token  *TokenDetails
	)

	err = c.ShouldBindJSON(&user)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	userID, err = h.useCases.SignIn(ctx, user.Username, user.Password)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	//call signin/signup
	token, err = CreateToken(userID, h.cfg.App.SecretAccess, h.cfg.App.SecretRefresh)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}

	//save redis ?
	err = h.useCases.CreateSession(ctx, token.AccessUuid, token.RefreshUuid, userID)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}

	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}
	
	responseWithStatus(c, http.StatusOK, "success signin", "OK", tokens)

}

func (h *Handler) SignUp(c *gin.Context) {

	var (
		user   *models.User
		err    error
		lastID int
	)

	err = c.ShouldBindJSON(&user)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lastID, err = h.useCases.SignUp(ctx, user)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success signup", "OK", lastID)

	//parseToken()
	c.Redirect(302, "/signin")

}
