package router

import (
	"github.com/gin-gonic/gin"

	"fateAdmin/docs"
	jwtauth "fateAdmin/pkg/jwt_auth"
	systemRouter "fateAdmin/router/system"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitSysRouter(r *gin.Engine, authMiddleware *jwtauth.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")

	systemRouter.SysBaseRouter(g)

	// swagger；注意：生产环境可以注释掉
	sysSwaggerRouter(g)

	// 无需认证
	systemRouter.SysNoCheckRoleRouter(g)

	// 需要认证
	sysCheckRoleRouterInit(g, authMiddleware)

	return g
}

func sysSwaggerRouter(r *gin.RouterGroup) {
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func sysCheckRoleRouterInit(r *gin.RouterGroup, authMiddleware *jwtauth.GinJWTMiddleware) {
	r.POST("/login", authMiddleware.LoginHandler)

	// Refresh time can be longer than token timeout
	r.GET("/refresh_token", authMiddleware.RefreshHandler)

	v1 := r.Group("/api/v1")

	systemRouter.RegisterSysUserRouter(v1, authMiddleware)
}
