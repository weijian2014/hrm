package db

type RoleMenu struct {
	Id       uint64
	RoleId   uint64
	MenuId   uint64
	ParentId uint64 // 子菜单
}
