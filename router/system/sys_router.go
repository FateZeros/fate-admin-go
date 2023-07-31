package system

import (
	"fateAdmin/apis/system"
	"fateAdmin/handler"

	jwt "fateAdmin/pkg/jwt_auth"

	"github.com/gin-gonic/gin"
)

func SysBaseRouter(r *gin.RouterGroup) {
	r.GET("/ping", handler.Ping)
}

func SysNoCheckRoleRouter(r *gin.RouterGroup) {
	v1 := r.Group("/api/v1")

	// 获取验证码
	v1.GET("/getCaptcha", system.GenerateCaptchaHandler)
	// 注册
	v1.POST("/register", system.RegisterSysUser)
}

func RegisterSysUserRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	sysuser := v1.Group("/sysUser").Use(authMiddleware.MiddlewareFunc())
	{
		sysuser.POST("", system.InsertSysUser)
		sysuser.GET("/:userId", system.GetSysUser)
	}
}
