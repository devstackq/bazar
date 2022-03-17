package v1

import (
	"database/sql"

	"github.com/devstackq/bazar/internal/config"
	"github.com/devstackq/bazar/internal/gallery/repository"
	"github.com/devstackq/bazar/internal/gallery/usecase"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

//cfg *config.Config,
func SetMachineEndpoints( cfg *config.Config, db *sql.DB, logger *logrus.Logger, group *gin.RouterGroup) {
	
	machineRepos := repository.MachineReposInit(db)
	machineUseCases  := usecase.UseCasesInit(machineRepos)

	handler := NewHandler(machineUseCases, logger)

	//filter - query=param
	machine := group.Group("/machine")
	{
		machine.POST("", handler.CreateMachine)
		machine.GET("/:id", handler.GetMachineByID)
		// machine.PATCH("/:id", handler.UpdateMachine)
		// machine.DELETE("/:id", handler.DeleteMachineByID)
		machine.GET("", handler.GetListMachine)
	}

	category := group.Group("/category")
	{
		category.POST("", handler.CreateCategory)
	}

	search := group.Group("/search")
	{
		search.POST("/:key_word", handler.Search)
	}

	trans := group.Group("/trans")
	{
		trans.POST("", handler.CreateTrans)
	}
}