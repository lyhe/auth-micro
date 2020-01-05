package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lyhe/auth-micro/router/api/v1"
)

func Router() *gin.Engine {
	router := gin.New()

	api := router.Group("/api/v1")
	{
		api.POST("/user/registry", v1.RegistryUser)
		api.POST("/user/login", v1.UserLogin)
	}

	return router
}
