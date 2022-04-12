package v1

import (
	"database/sql"

	"github.com/devstackq/bazar/internal/config"
	"github.com/devstackq/bazar/internal/gallery/repository"
	"github.com/devstackq/bazar/internal/gallery/usecase"
	"github.com/devstackq/bazar/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// auth; machine - crud; filter/sort/search
// profile; photo
// cfg *config.Config

func SetGalleryEndpoints(cfg *config.Config, db *sql.DB, logger *logrus.Logger, group *gin.RouterGroup) {
	machineRepos := repository.MachineReposInit(db)
	machineUseCases := usecase.UseCasesInit(machineRepos)

	handler := NewHandler(machineUseCases, logger)

	machine := group.Group("/machine")
	{
		machine.POST("", middleware.AuthorizeJWT("accessx"), handler.CreateMachine)
		machine.GET("/:id", handler.GetMachineByID)
		machine.GET("", handler.GetListMachine)
		machine.GET("/user", middleware.AuthorizeJWT("accessx"), handler.GetListMachineByUserID) // or companyID cars ?
		// machine.PATCH("/:id", handler.UpdateMachine)
		// machine.DELETE("/:id", handler.DeleteMachineByID)
		///v1/machine/user/:id :GET, user cereated cars
		machine.POST("/filter", handler.GetFilteredMachine)
		machine.POST("/upload/:id", middleware.AuthorizeJWT("accessx"), handler.Upload)
		machine.GET("/search", handler.Search)
	}
}
