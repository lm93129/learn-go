package xorm

// 创建数据库模型
type User struct {
	Id             int64  //如果无任何的tag，则默认为主键且自增
	UserName       string `xorm:"unique varchar(25) notnull"` // 用户名
	PasswordDigest string // 密码
	Email          string `xorm:"unique"` // 姓名
	Status         int    // 状态，1为启用，2为冻结
	Role           int    // 权限
	Avatar         string `xorm:"varchar(1024)"` //头像
}

type Permission struct {
	Name string `xorm:"unique"` //权限名称
	Id   int64
}

type RolesPermissions struct {
	RoleId       int64
	PermissionId int64
}

type Role struct {
	Id   int64
	Name string
}
