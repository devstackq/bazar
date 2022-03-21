package v1

import (
	"context"
	"log"
	"net/http"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

//todo: - token expires & user_id in redis -> check token_uuid;
//refresh -> call -> if access expired
//logout || refresh expired -> redirect signin

//client -> signup/signin -> createToken, save Db, uuid
//client -> send At; server Valid(at) - middleware
//server(at)- invalid -> client send - (refresh) -> create New At, Rt
//refresh - invalid -> redirect signin;
//logout -> removeSession db, client;

func (h *Handler) SignIn(c *gin.Context) {

	var (
		user   *models.User
		err    error
		userID int
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

	userID, err = h.useCases.SignIn(ctx, user.Username, user.Password)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}

	//call signin/signup/refresh
	token, err = CreateToken(userID, h.cfg.App.SecretAccess, h.cfg.App.SecretRefresh)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	token.UserID = userID

	// err = h.useCases.CreateSession(ctx, token)
	//refresh, /signin call
	// err = h.useCases.UpdateSession(ctx, token)

	// c.SetCookie("access-token", token.AccessToken, int(token.AtExpires), "", "", false, true)

	c.Writer.Header().Set("refresh_token", token.RefreshToken)
	c.Writer.Header().Set("access_token", token.AccessToken)

	//check - passwoir, email - from decrypt( passwoir, email) ; -> compare DB
	responseWithStatus(c, http.StatusOK, "success signin", "OK", token)
}

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

	lastID, err = h.useCases.SignUp(ctx, user)
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

//refresh - invalid -> redirect signin; else newToken
func (h *Handler) RefreshJwt(c *gin.Context) {
	//get refresh token; update AT, RT; - updateSessionDb();
	//if !Valid(refresh) -> redirect signin, 401
	//return token, 200
	//middleware -> refreshJwt, ->   err = h.useCases.CreateSession(ctx, token)

	var (
		err   error
		token *models.TokenDetails
	)

	userID, ok := c.Get("user_id")
	if !ok {
		log.Println("!have value")
		responseWithStatus(c, http.StatusBadRequest, "refresh token incorrect", "no have userI_id, to refresh token", nil)
		return
	}

	token, err = CreateToken(int(userID.(float64)), h.cfg.App.SecretAccess, h.cfg.App.SecretRefresh)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	// log.Println("set new token ", token.AccessToken)

	c.Writer.Header().Set("refresh_token", token.RefreshToken)
	c.Writer.Header().Set("access_token", token.AccessToken)

	responseWithStatus(c, http.StatusCreated, "refresh token created", "", nil)

}
check fundtional all
//logout -> removeSession db/redis, client;
func (h *Handler) Logout(c *gin.Context) {
todo;
	//remove client & db jwt token; by user_id
}
