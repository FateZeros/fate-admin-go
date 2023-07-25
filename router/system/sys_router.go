package system

import (
	"fateAdmin/apis/system"
	"fateAdmin/handler"

	"github.com/gin-gonic/gin"
)

func SysBaseRouter(r *gin.RouterGroup) {
	r.GET("/ping", handler.Ping)
}

func SysNoCheckRoleRouter(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")

	v1.GET("/getCaptcha", system.GenerateCaptchaHandler)
}
