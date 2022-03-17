package v1

import (
	"net/http"

	"github.com/devstackq/bazar/internal/gallery/models"
	"github.com/gin-gonic/gin"
)
func (h *Handler) CreateCategory(c *gin.Context) {

	var (
		cat models.Category		
		err    error
		lastID int
	)
	err = c.ShouldBindJSON(&cat)
	if err != nil {
		// h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	// if err = cat.Validate(); err != nil {
	lastID, err = h.useCases.CategoryUseCase.CreateCategory(&cat)
	if err != nil {
		// h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "category success created", "OK", lastID)
}