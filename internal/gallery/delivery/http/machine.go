package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

// CreateCar godoc
// @Description create car, header Authorization : access_token, body : {model.Machine}
// @Tags         Create car
// @Produce      json
// @Accept      json
// @Security BearerAuth
// @Param        input body  models.Machine true "model Machine"
// @Failure      400,500  {object}  models.Response
// @Success      200      integer  1
// @Router       /v1/machine [post]
func (h *Handler) CreateMachine(c *gin.Context) {
	var (
		machine models.Machine
		err     error
		lastID  int
	)
	// get user id from token
	userId, ok := c.Get("user_id")
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
	// ref int to float64
	machine.Creator.ID = strconv.Itoa(int(userId.(float64)))

	lastID, err = h.useCases.Create(&machine)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	// upload photo
	// c.Params = append(c.Params, gin.Param{"id", strconv.Itoa(lastID)})
	// h.Upload(c)

	responseWithStatus(c, http.StatusOK, "machine success created", "OK", lastID)
}

// GetCarByID godoc
// @Description GetCarByID path  : idCar
// @Tags         GetCarByID
// @Produce      json
// @Param        input path  string true "get car id from path"
// @Failure      400,500  {object}  models.Response
// @Success      200      {object} models.Machine
// @Router       /v1/machine/:id [get]
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

	result, err = h.useCases.GetMachineByID(id)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}

	result.Images, err = h.useCases.FileManagerUseCaseInterface.GetListSrc(id)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}

	responseWithStatus(c, http.StatusOK, "success retrun machine", "OK", result)
}

// GetListCar godoc
// @Description just  return ListCar by created_at - desc
// @Tags         GetListCar
// @Produce      json
// @Failure      400,500  {object}  models.Response
// @Success      200      {object} []models.Machine
// @Router       /v1/machine [get]
func (h *Handler) GetListMachine(c *gin.Context) {
	var (
		result  []*models.Machine
		err     error
		pageNum int
		value   string
	)
	// todo: DRY
	if value = c.Query("page_num"); value == "" {
		value = "1"
	}
	pageNum, err = strconv.Atoi(value)
	if err != nil || pageNum < 0 {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	result, err = h.useCases.GetRelevantMachines(pageNum) // default created date; desc
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	if len(result) == 0 {
		responseWithStatus(c, http.StatusNoContent, "now, empty machines", "OK", result)
		return
	}

	responseWithStatus(c, http.StatusOK, "success retrun list machines", "OK", result)
}

// GetListMachineByUserID godoc
// @Description GetListMachineByUserID , header Authorization : access_token, with query param, page_num=1
// @Tags         Machine
// @Produce      json
// @Security BearerAuth
// @Param        input query string true "with query param, page_num=1"
// @Failure      400,500  {object}  models.Response
// @Success      200     {object}  []models.Machine
// @Router       /v1/machine/user [get]
func (h *Handler) GetListMachineByUserID(c *gin.Context) {
	var (
		result  []*models.Machine
		err     error
		pageNum int
		value   string
	)
	userId, ok := c.Get("user_id")
	if !ok {
		h.logger.Info("no set user_id in context")
		return
	}
	if value = c.Query("page_num"); value == "" {
		value = "1"
	}
	pageNum, err = strconv.Atoi(value)
	if err != nil || pageNum < 0 {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	result, err = h.useCases.GetListMachineByUserID(userId.(float64), pageNum)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}

	if len(result) < 1 {
		responseWithStatus(c, http.StatusNoContent, "now, empty user created machines", "OK", result)
		return
	}

	responseWithStatus(c, http.StatusOK, "success returun user  list machines", "OK", result)
}
