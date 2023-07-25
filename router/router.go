package router

import (
	"github.com/gin-gonic/gin"

	jwtauth "fateAdmin/pkg/jwt_auth"
	systemRouter "fateAdmin/router/system"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwtauth.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")

	systemRouter.SysBaseRouter(g)

	// swagger；注意：生产环境可以注释掉
	// sysSwaggerRouter(g)

	// 无需认证
	systemRouter.SysNoCheckRoleRouter(g)

	// 需要认证
	//sysCheckRoleRouterInit(g, authMiddleware)

	return g
}
