package db

import "time"

type User struct {
	Id        uint64    `json:"id" description:"用户ID"`
	Name      string    `json:"username" description:"用户名"`
	Password  string    `json:"password" description:"密码"`
	Data      string    `json:"data" description:"用户数据"`
	UpdatedAt time.Time `json:"update_at" description:"更新时间"`
}

type Role struct {
	Id        uint64    `json:"id" description:"角色ID"`
	Name      string    `json:"name" description:"角色名"`
	UpdatedAt time.Time `json:"update_at" description:"更新时间"`
}

type Menu struct {
	Id        uint64    `json:"id" description:"菜单ID"`
	Name      string    `json:"name" description:"菜单名"`
	Url       string    `json:"url" description:"菜单链接"`
	ParentId  uint64    `json:"parent_id" description:"菜单的父级菜单ID"`
	UpdatedAt time.Time `json:"update_at" description:"更新时间"`
}

type UserRole struct {
	Id     uint64
	UserId uint64
	RoleId uint64
}

type RoleMenu struct {
	Id       uint64
	RoleId   uint64
	MenuId   uint64
	ParentId uint64 // 子菜单
}
