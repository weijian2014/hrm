package db

import (
	"time"

	"gorm.io/gorm"
)

type RoleRequest struct {
}

type Role struct {
	Id        uint64    `json:"id" description:"角色ID"`
	Name      string    `json:"name" description:"角色名"`
	UpdatedAt time.Time `json:"updated_at" description:"更新时间"`
}

func (r *Role) BeforeDelete(tx *gorm.DB) error {
	ur := &UserRole{
		RoleId: r.Id,
	}
	err := Delete(ur, "role_id = ?", ur.RoleId)
	if err != nil {
		return err
	}

	rm := &RoleMenu{
		RoleId: r.Id,
	}
	err = Delete(rm, "role_id = ?", rm.RoleId)
	if err != nil {
		return err
	}

	return nil
}
