package router

import (
	"fateAdmin/middleware"

	"github.com/gin-gonic/gin"
	// toolConfig "fateAdmin/tools/config"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	// if toolConfig.ApplicationConfig.IsHttps {
	// 	r.Use(handler)
	// }
	middleware.InitMiddleware(r)

	return r
}
