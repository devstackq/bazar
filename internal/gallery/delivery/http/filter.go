package v1

import (
	"net/http"

	"github.com/devstackq/bazar/internal/gallery/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetFilteredMachine(c *gin.Context) {

	var (
		result []*models.Machine
		err    error
	)
filterKeys := []string{"category", "state", "brand", "model"}

	keys, err := 	prepareQueryParam(c, filterKeys)

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
	
	responseWithStatus(c, http.StatusOK, "success return filtered cars", "OK", result)
}

// add model, brand, categoruy, state;
// add new cars; 
// check - filter~ 