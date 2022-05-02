package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCity(c *gin.Context) {
	var (
		city      *models.City
		err       error
		lastID    int
		countryID int
	)

	countryID, err = strconv.Atoi(c.Param("country_id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	err = c.ShouldBindJSON(&city)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	city.ID = countryID

	lastID, err = h.useCases.CityUseCaseInterface.CreateCity(city)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create City", "OK", lastID)
}

// GetListCityByCountryID godoc
// @Description GetListCityByCountryID return list cities by country id, Russia -> Moscow, Kazan, etc
// @Tags         Add-on list
// @Produce      json
// @Param        input path  integer true "/v1/city/country/:id"
// @Failure      400,500  {object}  models.Response
// @Success      200      {object} []models.City
// @Router       /v1/city/country/1 [get]
func (h *Handler) GetListCityByCountryID(c *gin.Context) {
	var (
		result    []*models.City
		err       error
		countryID int
	)

	countryID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	result, err = h.useCases.CityUseCaseInterface.GetListCityByCountryID(countryID)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list City", "OK", result)
}

func (h *Handler) GetCityByID(c *gin.Context) {
	var (
		result *models.City
		err    error
		id     int
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	result, err = h.useCases.CityUseCaseInterface.GetCityByID(id)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return City ", "OK", result)
}
