package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/gallery/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCity(c *gin.Context) {

	var (
		argument *models.City		
		err    error
		lastID int
	)

	err = c.ShouldBindJSON(&argument)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
		lastID, err = h.useCases.CityUseCaseInterface.CreateCity(argument)
		if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create City", "OK", lastID)
}

func (h *Handler) GetListCity(c *gin.Context) {

	var (
		result []*models.City		
		err    error
	)

	result, err = h.useCases.CityUseCaseInterface.GetListCity()
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
		id int
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


