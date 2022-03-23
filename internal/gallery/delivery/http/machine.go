package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateMachine(c *gin.Context) {
	var (
		machine models.Machine
		err     error
		lastID  int
	)

	uid, ok := c.Get("user_id")
	if !ok {
		h.logger.Info("no set user_id in context")
		return
	}

	err = c.ShouldBindJSON(&machine)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	machine.Creator.ID = strconv.Itoa(int(uid.(float64)))

	lastID, err = h.useCases.Create(&machine)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "machine success created", "OK", lastID)
}

func (h *Handler) GetMachineByID(c *gin.Context) {
	var (
		result *models.Machine
		err    error
		id     int
	)

	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}
	// getUserID; byCarID

	result, err = h.useCases.GetMachineByID(id)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success retrun machine", "OK", result)
}

func (h *Handler) GetListMachine(c *gin.Context) {
	var (
		result []*models.Machine
		err    error
	)
	result, err = h.useCases.GetRelevantMachines()
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success retrun list machines", "OK", result)
}

func (h *Handler) GetListMachineByUserID(c *gin.Context) {
	var (
		result []*models.Machine
		err    error
		id     int
	)
	id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}
	result, err = h.useCases.GetListMachineByUserID(id)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success returun user  list machines", "OK", result)
}
