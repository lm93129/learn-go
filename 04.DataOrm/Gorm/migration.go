package Gorm

func migration() {
	// 自动迁移模式
	DB.AutoMigrate(
		&User{},
		&Permission{},
		&Role{},
		&RolesPermissions{},
	)
	// 建立数据库外键
	DB.Model(&Permission{}).AddForeignKey("id", "user(id)",
		"CASCADE", "CASCADE")
	DB.Model(&RolesPermissions{}).AddForeignKey("role_id", "role(id)",
		"CASCADE", "CASCADE")
}
