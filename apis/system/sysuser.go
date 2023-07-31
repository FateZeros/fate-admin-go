package system

import (
	"fateAdmin/models/system"
	"fateAdmin/tools"
	"fateAdmin/tools/app"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary 注册用户
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body system.SysUser true "用户数据"
// @Success 200 {string} string	"{"code": 200, "message": "注册成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "注册失败"}"
// @Router /api/v1/register [post]
func RegisterSysUser(c *gin.Context) {
	var sysuser system.SysUser
	err := c.MustBindWith(&sysuser, binding.JSON)

	// 注册默认状态 和 角色 id 普通角色
	sysuser.Status = "0"
	sysuser.RoleId = 2

	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	id, err := sysuser.Insert()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}
	app.OK(c, id, "注册成功")
}

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

// @Summary 获取用户
// @Description 获取JSON
// @Tags 用户
// @Param userId path int true "用户编码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sysUser/{userId} [get]
// @Security Bearer
func GetSysUser(c *gin.Context) {
	var SysUser system.SysUser
	SysUser.UserId, _ = tools.StringToInt(c.Param("userId"))
	user, err := SysUser.Get()
	if err != nil {
		app.Error(c, -1, err, "")
		return
	}

	var SysRole system.SysRole

	roles, _ := SysRole.GetList()

	roleIds := make([]int, 0)
	roleIds = append(roleIds, user.RoleId)

	app.Custum(c, gin.H{
		"code":    200,
		"data":    user,
		"roleIds": roleIds,
		"roles":   roles,
	})
}
