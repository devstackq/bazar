package v1

import (
	"github.com/devstackq/bazar/internal/gallery/models"
	"github.com/gin-gonic/gin"
)

func responseWithStatus(c *gin.Context, status int, message, text string, data interface{}) {
	c.AbortWithStatusJSON(status, models.Response{
		Status:  text,
		Message: message,
		Data:    data,
	})
}