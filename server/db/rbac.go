package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Password string
	Data     string
}

type Role struct {
	gorm.Model
	Name string
}

type Menu struct {
	gorm.Model
	Name string
	Url  string
}

type UserRole struct {
	gorm.Model
	UserId uint
	RoleId uint
}

type RoleMenu struct {
	gorm.Model
	RoleId   uint
	MenuId   uint
	ParentId uint // 子菜单
}
