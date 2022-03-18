package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/models"
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
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	lastID, err = h.useCases.CategoryUseCaseInterface.CreateCategory(&cat)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "category success created", "OK", lastID)
}

func (h *Handler) GetListCategories(c *gin.Context) {

	var (
		result []*models.Category		
		err    error
	)

	result, err = h.useCases.CategoryUseCaseInterface.GetListCategories()
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list category", "OK", result)
}


func (h *Handler) GetCategoryByID(c *gin.Context) {
	var (
		result *models.Category		
		err    error
		id int
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	result, err = h.useCases.CategoryUseCaseInterface.GetByID(id)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return category ", "OK", result)
}