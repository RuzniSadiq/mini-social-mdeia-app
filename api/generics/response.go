package generics

import (
	"github.com/gin-gonic/gin"
)

func CustomJsonResponseData(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{"status_code": status,"data": data})
}

func CustomJsonResponseError(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"status_code": status,"error": message})
}
