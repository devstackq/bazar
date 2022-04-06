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

func prepareQueryParam(c *gin.Context, keys []string) (map[string]string, error) {
	result := map[string]string{}
	var value string

	for _, param := range keys {
		if value = c.Query(param); value != "" {
			result[param] = value
		}
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("filter param is empty")
	}
	return result, nil
}
