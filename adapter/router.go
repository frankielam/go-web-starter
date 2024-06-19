package adapter

import (
	"gowebstarter/infra/config"
	"log"
	"sync"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var (
	router        *gin.Engine
	ginOnce       sync.Once
	jwtMiddleware *jwt.GinJWTMiddleware
)

func InitGin() *gin.Engine {
	ginOnce.Do(func() {
		conf := config.GetConf()
		if len(conf.Server.Mode) > 0 {
			gin.SetMode(conf.Server.Mode)
		}

		router = gin.Default()
		initRouter(router)

		// JWT start
		jwtMiddleware, err := ginJWTMiddleware()
		if err != nil {
			panic(err)
			log.Fatalf("Failed to initialize JWT middleware: %v", err)
		}
		privateRouter(router, jwtMiddleware)
		publicRouter(router, jwtMiddleware)
		// JWT end

	})
	return router
}

func initRouter(r *gin.Engine) {
	gin.DisableConsoleColor()
	r.GET("/", index)
}

func publicRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	r.POST("/login", authMiddleware.LoginHandler)
	// router.Static("/static", "")
}

func privateRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	my := r.Group("/my")
	my.Use(authMiddleware.MiddlewareFunc())
	{
		my.GET("/profile", profile)
	}
}
