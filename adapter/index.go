package adapter

import (
	"gowebstarter/infra/db"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	user, err := db.GetSysUserById(c, 1)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, user)
}

func profile(c *gin.Context) {
	userId, _ := GetIdFromJWT(c)
	slog.Info("userId", userId)
	user, err := db.GetSysUserById(c, userId)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, user)
}
