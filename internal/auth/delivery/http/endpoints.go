package v1

import (
	"database/sql"

	"github.com/devstackq/bazar/internal/auth/repository/psql"
	"github.com/devstackq/bazar/internal/auth/usecase"
	"github.com/devstackq/bazar/internal/config"
	"github.com/devstackq/bazar/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetAuthEndpoints( cfg *config.Config, db *sql.DB, logger *logrus.Logger, group *gin.RouterGroup) {
	
	authRepo := psql.AuthRepositoryInit(db)
	authUseCases  := usecase.AuthUseCaseInit(authRepo, cfg.App.HashSalt, cfg.App.TokenTTL)
	
	handler := NewHandler(authUseCases, logger, cfg)

	auth := group.Group("/auth")
	{
		auth.POST("/signup", handler.SignUp)
		auth.POST("/signin", handler.SignIn)
	}

	refresh := group.Group("/auth/refresh", middleware.AuthorizeJWT("refreshx"))//todo: env config
	{
		refresh.POST("", handler.RefreshJwt)
	}

	logout := group.Group("/auth/logout", middleware.AuthorizeJWT("accessx"))
	{
		logout.POST("", handler.Logout)
	}
}