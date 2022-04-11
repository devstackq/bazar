package v1

import (
	"database/sql"

	"github.com/devstackq/bazar/internal/admin/repository"
	"github.com/devstackq/bazar/internal/admin/usecase"
	"github.com/devstackq/bazar/internal/config"
	"github.com/devstackq/bazar/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// update further

func SetAdminEndpoints(cfg *config.Config, db *sql.DB, logger *logrus.Logger, group *gin.RouterGroup) {
	adminRepos := repository.AdminReposInit(db)
	adminUseCases := usecase.UseCasesInit(adminRepos)

	handler := NewHandler(adminUseCases, logger)

	role := group.Group("/role", middleware.AuthorizeJWT("accessx"))
	{
		role.POST("", handler.CreateRole)
		role.GET("", handler.GetListRole)

	}

	category := group.Group("/category")
	{
		category.POST("", middleware.AuthorizeJWT("accessx"), handler.CreateCategory)
		category.GET("", handler.GetListCategories)
		category.GET("/:id", handler.GetCategoryByID)
	}

	trans := group.Group("/transmission")
	{
		trans.POST("", middleware.AuthorizeJWT("accessx"), handler.CreateTransmission)
		trans.GET("", handler.GetListTransmission)
		trans.GET("/:id", handler.GetTransmissionByID)
	}

	state := group.Group("/state")
	{
		state.POST("", middleware.AuthorizeJWT("accessx"), handler.CreateState)
		state.GET("", handler.GetListState)
		state.GET("/:id", handler.GetStateByID)
	}

	country := group.Group("/country")
	{
		country.POST("", middleware.AuthorizeJWT("accessx"), handler.CreateCountry)
		country.GET("", handler.GetListCountry)
		country.GET("/:id", handler.GetCountryByID)
	}

	city := group.Group("/city")
	{
		city.POST("/:country_id", middleware.AuthorizeJWT("accessx"), handler.CreateCity)
		city.GET("/country/:id", handler.GetListCityByCountryID)
		city.GET("/:id", handler.GetCityByID)
	}

	brand := group.Group("/brand")
	{
		brand.POST("", middleware.AuthorizeJWT("accessx"), handler.CreateBrand)
		brand.GET("", handler.GetListBrand)
		brand.GET("/:id", handler.GetBrandByID)
	}
	model := group.Group("/model")
	{
		model.POST("/:model_id", middleware.AuthorizeJWT("accessx"), handler.CreateModel)
		model.GET("/brand/:id", handler.GetListModelByBrandID)
		model.GET("/:id", handler.GetModelByID)
	}
	fuel := group.Group("/fuel")
	{
		fuel.POST("", middleware.AuthorizeJWT("accessx"), handler.CreateFuel)
		fuel.GET("", handler.GetListFuel)
		fuel.GET("/:id", handler.GetFuelByID)
	}
	driveUnit := group.Group("/drive_unit")
	{
		driveUnit.POST("", middleware.AuthorizeJWT("accessx"), handler.CreateDriveUnit)
		driveUnit.GET("", handler.GetListDriveUnit)
		driveUnit.GET("/:id", handler.GetDriveUnitByID)
	}
	bodyType := group.Group("/body_type")
	{
		bodyType.POST("", middleware.AuthorizeJWT("accessx"), handler.CreateBodyType)
		bodyType.GET("", handler.GetListBodyType)
		bodyType.GET("/:id", handler.GetBodyTypeByID)
	}

	color := group.Group("/color")
	{
		color.POST("", middleware.AuthorizeJWT("accessx"), handler.CreateColor)
		color.GET("", handler.GetListColor)
		color.GET("/:id", handler.GetColorByID)
	}
}
