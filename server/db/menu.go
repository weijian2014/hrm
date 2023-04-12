package db

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	Id        uint64    `json:"id" description:"菜单ID"`
	Name      string    `json:"name" description:"菜单名"`
	ParentId  uint64    `json:"parent_id" description:"菜单的父级菜单ID"`
	Icon      string    `json:"icon" description:"菜单图标"`
	Url       string    `json:"url" description:"菜单链接"`
	UpdatedAt time.Time `json:"updated_at" description:"更新时间"`
}

// Hook是事务的, 返回错误事务将终止, 并执行回滚
func (m *Menu) BeforeDelete(tx *gorm.DB) error {
	rm := &RoleMenu{
		MenuId: m.Id,
	}

	err := Delete(rm, "menu_id = ?", rm.MenuId)
	if err != nil {
		return err
	}
	return nil
}
