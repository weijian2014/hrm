package db

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	Id        uint64    `json:"id" description:"菜单ID"`
	Name      string    `json:"name" description:"菜单名"`
	Url       string    `json:"url" description:"菜单链接"`
	ParentId  uint64    `json:"parent_id" description:"菜单的父级菜单ID"`
	UpdatedAt time.Time `json:"updated_at" description:"更新时间"`
}

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
