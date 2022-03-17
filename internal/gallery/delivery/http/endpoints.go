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
		machine.GET("", handler.GetListMachine)
		machine.GET("/user/:id", handler.GetListMachineByUserID)
		// machine.PATCH("/:id", handler.UpdateMachine)
		// machine.DELETE("/:id", handler.DeleteMachineByID)
	}

	search := group.Group("/search")
	{
		search.POST("/:key_word", handler.Search)
	}

	category := group.Group("/category")
	{
		category.POST("", handler.CreateCategory)
		category.GET("", handler.GetListCategories)
		category.GET("/:id", handler.GetCategoryByID)
	}

	trans := group.Group("/trans")
	{
		trans.POST("", handler.CreateTransmission)
		trans.GET("", handler.GetListTransmission)
		trans.GET("/:id", handler.GetTransmissionByID)
	}

	country := group.Group("/country")
	{
		country.POST("", handler.CreateCountry)
		country.GET("", handler.GetListCountry)
		country.GET("/:id", handler.GetCountryByID)
	}

	city := group.Group("/city")
	{
		city.POST("", handler.CreateCity)
		city.GET("", handler.GetListCity)
		city.GET("/:id", handler.GetCityByID)
	}

	state := group.Group("/state")
	{
		state.POST("", handler.CreateState)
		state.GET("", handler.GetListState)
		state.GET("/:id", handler.GetStateByID)
	}
	brand := group.Group("/brand")
	{
		brand.POST("", handler.CreateBrand)
		brand.GET("", handler.GetListBrand)
		brand.GET("/:id", handler.GetBrandByID)
	}
	model := group.Group("/model")
	{
		model.POST("", handler.CreateModel)
		model.GET("", handler.GetListModel)
		model.GET("/:id", handler.GetModelByID)
	}
	fuel := group.Group("/fuel")
	{
		fuel.POST("", handler.CreateFuel)
		fuel.GET("", handler.GetListFuel)
		fuel.GET("/:id", handler.GetFuelByID)
	}
	driveUnit := group.Group("/drive_unit")
	{
		driveUnit.POST("", handler.CreateDriveUnit)
		driveUnit.GET("", handler.GetListDriveUnit)
		driveUnit.GET("/:id", handler.GetDriveUnitByID)
	}
	bodyType := group.Group("/body_type")
	{
		bodyType.POST("", handler.CreateBodyType)
		bodyType.GET("", handler.GetListBodyType)
		bodyType.GET("/:id", handler.GetBodyTypeByID)
	}

	color := group.Group("/color")
	{
		color.POST("", handler.CreateColor)
		color.GET("", handler.GetListColor)
		color.GET("/:id", handler.GetColorByID)
	}
}

//todo: GetList, GetByID, Create; add all realation table
//start file server 