package tools

import (
	jwt "fateAdmin/pkg/jwt_auth"
	"fateAdmin/pkg/logger"

	"github.com/gin-gonic/gin"
)

func ExtractClaims(c *gin.Context) jwt.MapClaims {
	claims, exists := c.Get("JWT_PAYLOAD")
	if !exists {
		return make(jwt.MapClaims)
	}

	return claims.(jwt.MapClaims)
}

func GetUserIdStr(c *gin.Context) string {
	data := ExtractClaims(c)
	if data["identity"] != nil {
		return Int64ToString(int64((data["identity"]).(float64)))
	}
	logger.Info("********** 路径：" + c.Request.URL.Path + "  请求方法：" + c.Request.Method + "  缺少identity")
	return ""
}
