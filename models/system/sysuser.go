package system

import (
	"errors"
	"fateAdmin/global/orm"

	"golang.org/x/crypto/bcrypt"
)

type SysUserId struct {
	UserId int `gorm:"primary_key;AUTO_INCREMENT"  json:"userId"` // 编码
}

type SysUserB struct {
	NickName string `gorm:"type:varchar(128)" json:"nickName"` // 昵称
	Phone    string `gorm:"type:varchar(11)" json:"phone"`     // 手机号
	RoleId   int    `gorm:"type:int(11)" json:"roleId"`        // 角色编码
	Salt     string `gorm:"type:varchar(255)" json:"salt"`     //盐
	Avatar   string `gorm:"type:varchar(255)" json:"avatar"`   //头像
	Sex      string `gorm:"type:varchar(255)" json:"sex"`      //性别
	Email    string `gorm:"type:varchar(128)" json:"email"`    //邮箱
	DeptId   int    `gorm:"type:int(11)" json:"deptId"`        //部门编码
	PostId   int    `gorm:"type:int(11)" json:"postId"`        //职位编码
	CreateBy string `gorm:"type:varchar(128)" json:"createBy"` //
	UpdateBy string `gorm:"type:varchar(128)" json:"updateBy"` //
	Remark   string `gorm:"type:varchar(255)" json:"remark"`   //备注
	Status   string `gorm:"type:int(1);" json:"status"`
	Params   string `gorm:"-" json:"params"`
	BaseModel
}

type UserName struct {
	Username string `gorm:"type:varchar(64)" json:"username"`
}

type PassWord struct {
	// 密码
	Password string `gorm:"type:varchar(128)" json:"password"`
}

type LoginM struct {
	UserName
	PassWord
}

type SysUser struct {
	SysUserId
	SysUserB
	LoginM
}

func (SysUser) TableName() string {
	return "sys_user"
}

// 加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

// 新增用户
func (e SysUser) Insert() (id int, err error) {
	if err = e.Encrypt(); err != nil {
		return
	}

	// check 用户名
	var count int
	orm.Eloquent.Table(e.TableName()).Where("username = ? and `delete_time` IS NULL", e.Username).Count(&count)
	if count > 0 {
		err = errors.New("账户已存在！")
		return
	}

	//添加数据
	if err = orm.Eloquent.Table(e.TableName()).Create(&e).Error; err != nil {
		return
	}
	id = e.UserId
	return
}
