package v1

import (
	"net/http"

	"github.com/devstackq/bazar/internal/gallery/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTrans(c *gin.Context) {

	var (
		argument *models.Transmission		
		err    error
		lastID int
	)

	err = c.ShouldBindJSON(&argument)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
		lastID, err = h.useCases.TransUseCase.CreateTransmission(argument)
		if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create transmission", "OK", lastID)
}