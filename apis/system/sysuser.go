package system

import (
	"fateAdmin/models/system"
	"fateAdmin/tools"
	"fateAdmin/tools/app"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary 创建用户
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body system.SysUser true "用户数据"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/sysUser [post]
func InsertSysUser(c *gin.Context) {
	var sysuser system.SysUser
	err := c.MustBindWith(&sysuser, binding.JSON)
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	sysuser.CreateBy = tools.GetUserIdStr(c)
	id, err := sysuser.Insert()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, id, "添加成功")
}
