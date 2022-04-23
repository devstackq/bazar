package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetFilteredMachine(c *gin.Context) {
	var (
		result  []*models.Machine
		err     error
		value   string
		pageNum int
	)

	f := models.NewQueryParams()

	keys, err := prepareQueryParam(c, f)
	// keys, err := prepareQueryParam(c, filterKeys)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "query param error", nil)
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

	result, err = h.useCases.FilterUseCaseInterface.GetListMachineByFilter(keys, pageNum)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	if len(result) == 0 {
		responseWithStatus(c, http.StatusNotFound, "Info, not found filtered items", "Info:", nil)
		return
	}
	// log.Println(result[0].MainImage, result[1].MainImage, 123)
	// log.Println(result, " filtered vals")
	responseWithStatus(c, http.StatusOK, "success return filtered cars", "OK", result)
}