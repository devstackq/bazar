package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/gallery/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateFuel(c *gin.Context) {

	var (
		argument *models.Fuel		
		err    error
		lastID int
	)

	err = c.ShouldBindJSON(&argument)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
		lastID, err = h.useCases.FuelUseCaseInterface.CreateFuel(argument)
		if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create Fuel", "OK", lastID)
}

func (h *Handler) GetListFuel(c *gin.Context) {

	var (
		result []*models.Fuel		
		err    error
	)

	result, err = h.useCases.FuelUseCaseInterface.GetListFuel()
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list Fuel", "OK", result)
}

func (h *Handler) GetFuelByID(c *gin.Context) {
	var (
		result *models.Fuel		
		err    error
		id int
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	result, err = h.useCases.FuelUseCaseInterface.GetFuelByID(id)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return Fuel ", "OK", result)
}


