package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint64    `json:"id" description:"用户ID"`
	Name      string    `json:"username" description:"登录的用户名"`
	Password  string    `json:"password" description:"登录的密码"`
	Data      string    `json:"data" description:"用户数据"`
	CreatedAt time.Time `json:"created_at" description:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" description:"更新时间"`
}

// Hook是事务的, 返回错误事务将终止, 并执行回滚
func (u *User) BeforeDelete(tx *gorm.DB) error {
	ur := &UserRole{
		UserId: u.Id,
	}

	err := Delete(ur, "user_id = ?", ur.UserId)
	if err != nil {
		return err
	}
	return nil
}
