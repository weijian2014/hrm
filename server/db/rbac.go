package db

type User struct {
	Id       uint64
	Name     string
	Password string
	Data     string
}

type Role struct {
	Id   uint64
	Name string
}

type Menu struct {
	Id   uint64
	Name string
	Url  string
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
