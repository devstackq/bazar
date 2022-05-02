package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateModel(c *gin.Context) {
	var (
		model   *models.Model
		err     error
		lastID  int
		brandID int
	)
	brandID, err = strconv.Atoi(c.Param("model_id"))
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

// GetListModelByBrandID godoc
// @Description GetListModelByBrandID return list model by brandId, Tesla -> ModelX, S3
// @Tags         Add-on list
// @Produce      json
// @Param        input path  string true "/v1/model/brand/:id"
// @Failure      400,500  {object}  models.Response
// @Success      200      {object} []models.Model
// @Router       /v1/model/brand/1 [get]
func (h *Handler) GetListModelByBrandID(c *gin.Context) {
	var (
		result  []*models.Model
		err     error
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
		id     int
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
