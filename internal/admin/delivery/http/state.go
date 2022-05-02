package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateState(c *gin.Context) {
	var (
		argument *models.State
		err      error
		lastID   int
	)

	err = c.ShouldBindJSON(&argument)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	lastID, err = h.useCases.StateUseCaseInterface.CreateState(argument)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create State", "OK", lastID)
}

// GetListState godoc
// @Description GetListState return list state, new, second hand, crash
// @Tags         Add-on list
// @Produce      json
// @Failure      400,500  {object}  models.Response
// @Success      200      {object} []models.State
// @Router       /v1/state [get]
func (h *Handler) GetListState(c *gin.Context) {
	var (
		result []*models.State
		err    error
	)

	result, err = h.useCases.StateUseCaseInterface.GetListState()
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list State", "OK", result)
}

func (h *Handler) GetStateByID(c *gin.Context) {
	var (
		result *models.State
		err    error
		id     int
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	result, err = h.useCases.StateUseCaseInterface.GetStateByID(id)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return State ", "OK", result)
}
