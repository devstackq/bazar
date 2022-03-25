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

// cfg *config.Config,
func SetAdminEndpoints(cfg *config.Config, db *sql.DB, logger *logrus.Logger, group *gin.RouterGroup) {
	adminRepos := repository.AdminReposInit(db)
	adminUseCases := usecase.UseCasesInit(adminRepos)

	handler := NewHandler(adminUseCases, logger)

	role := group.Group("/role", middleware.AuthorizeJWT("accessx"))
	{
		role.POST("", handler.CreateRole)
		role.GET("", handler.GetListRole)

	}

	category := group.Group("/category", middleware.AuthorizeJWT("accessx"))
	{
		category.POST("", handler.CreateCategory)
		category.GET("", handler.GetListCategories)
		category.GET("/:id", handler.GetCategoryByID)
	}

	trans := group.Group("/trans", middleware.AuthorizeJWT("accessx"))
	{
		trans.POST("", handler.CreateTransmission)
		trans.GET("", handler.GetListTransmission)
		trans.GET("/:id", handler.GetTransmissionByID)
	}
	state := group.Group("/state", middleware.AuthorizeJWT("accessx"))
	{
		state.POST("", handler.CreateState)
		state.GET("", handler.GetListState)
		state.GET("/:id", handler.GetStateByID)
	}

	country := group.Group("/country", middleware.AuthorizeJWT("accessx"))
	{
		country.POST("", handler.CreateCountry)
		country.GET("", handler.GetListCountry)
		country.GET("/:id", handler.GetCountryByID)
	}

	city := group.Group("/city", middleware.AuthorizeJWT("accessx"))
	{
		city.POST("/:country_id", handler.CreateCity)
		city.GET("country/:id", handler.GetListCityByCountryID)
		city.GET("/:id", handler.GetCityByID)
	}
	brand := group.Group("/brand", middleware.AuthorizeJWT("accessx"))
	{
		brand.POST("", handler.CreateBrand)
		brand.GET("", handler.GetListBrand)
		brand.GET("/:id", handler.GetBrandByID)
	}
	model := group.Group("/model", middleware.AuthorizeJWT("accessx"))
	{
		model.POST("/:model_id", handler.CreateModel)
		model.GET("/brand/:id", handler.GetListModelByBrandID)
		model.GET("/:id", handler.GetModelByID)
	}
	fuel := group.Group("/fuel", middleware.AuthorizeJWT("accessx"))
	{
		fuel.POST("", handler.CreateFuel)
		fuel.GET("", handler.GetListFuel)
		fuel.GET("/:id", handler.GetFuelByID)
	}
	driveUnit := group.Group("/drive_unit", middleware.AuthorizeJWT("accessx"))
	{
		driveUnit.POST("", handler.CreateDriveUnit)
		driveUnit.GET("", handler.GetListDriveUnit)
		driveUnit.GET("/:id", handler.GetDriveUnitByID)
	}
	bodyType := group.Group("/body_type", middleware.AuthorizeJWT("accessx"))
	{
		bodyType.POST("", handler.CreateBodyType)
		bodyType.GET("", handler.GetListBodyType)
		bodyType.GET("/:id", handler.GetBodyTypeByID)
	}

	color := group.Group("/color", middleware.AuthorizeJWT("accessx"))
	{
		color.POST("", handler.CreateColor)
		color.GET("", handler.GetListColor)
		color.GET("/:id", handler.GetColorByID)
	}
}
