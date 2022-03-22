package v1

import (
	"database/sql"

	"github.com/devstackq/bazar/internal/config"
	"github.com/devstackq/bazar/internal/middleware"
	"github.com/devstackq/bazar/internal/profile/repository/psql"
	"github.com/devstackq/bazar/internal/profile/usecase"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetprofileEndpoints( cfg *config.Config, db *sql.DB, logger *logrus.Logger, group *gin.RouterGroup) {
	
	profileRepo := psql.ProfileRepositoryInit(db)
	profileUseCase  := usecase.ProfileUseCaseInit(profileRepo)

	handler := NewHandler(profileUseCase, logger, cfg)

	profile := group.Group("/profile", middleware.AuthorizeJWT("accessx"))
	{
		profile.GET("", handler.GetProfileBio)
	}
}