package v1

import (
	"net/http"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetListRole(c *gin.Context) {
	var (
		result []*models.Role
		err    error
	)

	result, err = h.useCases.RoleUseCaseInterface.GetListRole()
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list Roles", "OK", result)
}

func (h *Handler) CreateRole(c *gin.Context) {
	var (
		role   *models.Role
		err    error
		lastID int
	)

	err = c.ShouldBindJSON(&role)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}

	lastID, err = h.useCases.RoleUseCaseInterface.CreateRole(role)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create user Role", "OK", lastID)
}
