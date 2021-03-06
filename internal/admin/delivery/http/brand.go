package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateBrand(c *gin.Context) {
	var (
		brand  *models.Brand
		err    error
		lastID int
	)

	err = c.ShouldBindJSON(&brand)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	lastID, err = h.useCases.BrandUseCaseInterface.CreateBrand(brand)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create Brand", "OK", lastID)
}

// GetListBrand godoc
// @Description GetListBrand return list car brand, tesla, vaz, hummer
// @Tags         Add-on list
// @Produce      json
// @Failure      400,500  {object}  models.Response
// @Success      200      {object} []models.Brand
// @Router       /v1/brand [get]
func (h *Handler) GetListBrand(c *gin.Context) {
	var (
		result []*models.Brand
		err    error
	)

	result, err = h.useCases.BrandUseCaseInterface.GetListBrand()
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list Brand", "OK", result)
}

func (h *Handler) GetBrandByID(c *gin.Context) {
	var (
		result *models.Brand
		err    error
		id     int
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	result, err = h.useCases.BrandUseCaseInterface.GetBrandByID(id)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return Brand ", "OK", result)
}
