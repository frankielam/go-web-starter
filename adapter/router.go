package adapter

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	router  *gin.Engine
	ginOnce sync.Once
)

func InitGin() *gin.Engine {
	ginOnce.Do(func() {
		router = gin.Default()
		initRouter(router)

	})
	return router
}

func initRouter(r *gin.Engine) {
	gin.DisableConsoleColor()
	r.GET("/", index)

}

func publicRouter(r *gin.Engine) {
	r.Static("/static", "")
}
