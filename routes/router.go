package routes

import (
	apiV1 "github.com/gethinyan/sso-check/api/v1"
	"github.com/gethinyan/sso-check/middleware/jwt"
	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.Use(jwt.JWT())
		{
			v1.GET("/loginCheck", apiV1.LoginCheck)
		}
	}

	return router
}
