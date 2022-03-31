package v1

import (
	"log"
	"net/http"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetFilteredMachine(c *gin.Context) {
	var (
		result []*models.Machine
		err    error
	)
	// getDataFromDb
	filterKeys := []string{
		"category", "state", "brand",
		"model", "priceFrom", "priceTo",
		"yearFrom", "yearTo",
		"sort_created_at", "sort_price",
		"sort_year", "sort_odometer",
	}

	keys, err := prepareQueryParam(c, filterKeys)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "query param error", nil)
		return
	}

	result, err = h.useCases.FilterUseCaseInterface.GetListMachineByFilter(keys)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	if len(result) == 0 {
		responseWithStatus(c, http.StatusNotFound, "Info, not found filtered items", "Info:", nil)
		return
	}
	log.Println(result[0].MainImage, result[1].MainImage, 123)
	// log.Println(result, " filtered vals")
	responseWithStatus(c, http.StatusOK, "success return filtered cars", "OK", result)
}
