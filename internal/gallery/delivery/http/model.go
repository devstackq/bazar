package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/gallery/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateModel(c *gin.Context) {

	var (
		model *models.Model		
		err    error
		lastID int
		brandID int
	)
	brandID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	err = c.ShouldBindJSON(&model)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	
	model.ID = brandID

		lastID, err = h.useCases.ModelUseCaseInterface.CreateModel(model)
		if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create Model", "OK", lastID)
}

func (h *Handler) GetListModelByBrandID(c *gin.Context) {

	var (
		result []*models.Model		
		err    error
		brandID int
	)

	brandID, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	result, err = h.useCases.ModelUseCaseInterface.GetListModelByBrandID(brandID)

	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list Model", "OK", result)
}

func (h *Handler) GetModelByID(c *gin.Context) {
	var (
		result *models.Model		
		err    error
		id int
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	result, err = h.useCases.ModelUseCaseInterface.GetModelByID(id)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return Model ", "OK", result)
}


