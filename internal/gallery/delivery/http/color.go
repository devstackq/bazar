package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateColor(c *gin.Context) {

	var (
		argument *models.Color		
		err    error
		lastID int
	)

	err = c.ShouldBindJSON(&argument)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
		lastID, err = h.useCases.ColorUseCaseInterface.CreateColor(argument)
		if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create Color", "OK", lastID)
}

func (h *Handler) GetListColor(c *gin.Context) {

	var (
		result []*models.Color		
		err    error
	)

	result, err = h.useCases.ColorUseCaseInterface.GetListColor()
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list Color", "OK", result)
}

func (h *Handler) GetColorByID(c *gin.Context) {
	var (
		result *models.Color		
		err    error
		id int
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	result, err = h.useCases.ColorUseCaseInterface.GetColorByID(id)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return Color ", "OK", result)
}


