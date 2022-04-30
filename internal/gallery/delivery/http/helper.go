package v1

import (
	"fmt"

	"github.com/devstackq/bazar/internal/models"
	"github.com/gin-gonic/gin"
)

func responseWithStatus(c *gin.Context, status int, message, text string, data interface{}) {
	c.AbortWithStatusJSON(status, models.Response{
		Status:  text,
		Message: message,
		Data:    data,
	})
}

func prepareQueryParam(c *gin.Context, keys *models.QueryParams) (*models.QueryParams, error) {
	var value string
	// ref
	keys.Filter["category"] = ""
	keys.Filter["state"] = ""
	keys.Filter["brand"] = ""
	keys.Filter["model"] = ""
	keys.Filter["priceFrom"] = ""
	keys.Filter["priceTo"] = ""
	keys.Filter["yearFrom"] = ""
	keys.Filter["yearTo"] = ""

	keys.Sort["sort_created_at"] = ""
	keys.Sort["sort_price"] = ""
	keys.Sort["sort_year"] = ""
	keys.Sort["sort_odometer"] = ""

	for key := range keys.Filter {
		if value = c.Query(key); value != "" {
			keys.Filter[key] = value
		}
	}
	for key := range keys.Sort {
		if value = c.Query(key); value != "" {
			keys.Sort[key] = value
		}
	}
	if len(keys.Filter) == 0 && len(keys.Sort) == 0 {
		return nil, fmt.Errorf("Filter && sort params is empty")
	}
	return keys, nil
}
