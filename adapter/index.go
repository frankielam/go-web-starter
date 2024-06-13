package adapter

import (
	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, Gin!",
	})
}
