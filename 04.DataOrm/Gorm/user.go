package Gorm

import (
	"github.com/jinzhu/gorm"
)

// 创建数据库模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"` // 用户名
	PasswordDigest string // 密码
	Email          string `gorm:"unique"` // 姓名
	Status         int    // 状态，1为启用，2为冻结
	Role           int    // 权限
	Avatar         string `gorm:"size:1000"` //头像
}

type Permission struct {
	Name string `gorm:"unique"` //权限名称
	Id   uint   `gorm:"primary_key;not null;AUTO_INCREMENT"`
}

type RolesPermissions struct {
	RoleId       uint
	PermissionId uint
}

type Role struct {
	Id   uint
	Name string
}
