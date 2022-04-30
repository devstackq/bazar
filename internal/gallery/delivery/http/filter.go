package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

// Filter  godoc
// @Description  Get Filtered Cars, recieve by query-params ["category", "state", "brand", "model"] and/or [priceTo/proceFrom] and/or 1 param - sort:  [sort_created_at/sort_price/sort_year/sort_odometer - asc/desc]  and  page_num=1
// @Tags         Filter Cars
// @Produce      json
// @Param        input  query   string  true "?category=1&state=1&brand=1&model=1&priceFrom=1000&priceTo=20000&sort_price=asc&page_num=1"
// @Failure      400,500  {object}  models.Response
// @Success      200      {object}  []models.Machine
// @Router       /v1/machine/filter [get]
func (h *Handler) GetFilteredMachine(c *gin.Context) {
	var (
		result  []*models.Machine
		err     error
		value   string
		pageNum int
	)

	f := models.NewQueryParams()

	keys, err := prepareQueryParam(c, f)
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
	responseWithStatus(c, http.StatusOK, "success return filtered cars", "OK", result)
}
