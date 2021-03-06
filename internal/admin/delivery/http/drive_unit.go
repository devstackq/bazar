package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateDriveUnit(c *gin.Context) {
	var (
		argument *models.DriveUnit
		err      error
		lastID   int
	)

	err = c.ShouldBindJSON(&argument)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	lastID, err = h.useCases.DriveUnitUseCaseInterface.CreateDriveUnit(argument)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create DriveUnit", "OK", lastID)
}

// GetListDriveUnit godoc
// @Description GetListDriveUnit return list drive unit, up, back, full
// @Tags         Add-on list
// @Produce      json
// @Failure      400,500  {object}  models.Response
// @Success      200      {object} []models.DriveUnit
// @Router       /v1/drive_unit [get]
func (h *Handler) GetListDriveUnit(c *gin.Context) {
	var (
		result []*models.DriveUnit
		err    error
	)

	result, err = h.useCases.DriveUnitUseCaseInterface.GetListDriveUnit()
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list DriveUnit", "OK", result)
}

func (h *Handler) GetDriveUnitByID(c *gin.Context) {
	var (
		result *models.DriveUnit
		err    error
		id     int
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}
	result, err = h.useCases.DriveUnitUseCaseInterface.GetDriveUnitByID(id)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return DriveUnit ", "OK", result)
}
