package v1

import (
	"net/http"
	"strconv"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Search(c *gin.Context) {
	var (
		result  []*models.Machine
		err     error
		keyWord string
		value   string
		pageNum int
	)
	// keyWord = c.Param("key_word")

	if keyWord = c.Query("key_word"); keyWord == "" {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	if value = c.Query("page_num"); value == "" {
		//set default
		value = "1"
	}

	pageNum, err = strconv.Atoi(value)
	if err != nil || pageNum < 0 {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	result, err = h.useCases.SearchUseCaseInterface.SearchByKeyWord(keyWord, pageNum)

	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}

	if len(result) == 0 {
		responseWithStatus(c, http.StatusNotFound, "not found by keyword", "Info", result)
		return
	}
	responseWithStatus(c, http.StatusOK, "find machine by keyword", "OK", result)
}
