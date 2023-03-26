package db

import (
	"time"

	"gorm.io/gorm"
)

// func (u *User) CreateTable() error {
// 	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
// 	if err != nil {
// 		return err
// 	}

// 	// 表名为users
// 	if err = db.AutoMigrate(u); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (u *User) Find() error {
// 	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
// 	if err != nil {
// 		return err
// 	}

// 	ret := db.First(u, "name = ?", u.Name)
// 	if ret.Error != nil {
// 		return ret.Error
// 	}

// 	return nil
// }

// func (u *User) Insert() error {
// 	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
// 	if err != nil {
// 		return err
// 	}

// 	ret := db.Create(u)
// 	if ret.Error != nil {
// 		return ret.Error
// 	}

// 	return nil
// }

type UserLoginRequest struct {
	Name     string `json:"username" description:"登录的用户名"`
	Password string `json:"password" description:"登录的密码"`
}

type UserUpdateRequest struct {
	Id       uint64 `json:"id" description:"用户ID"`
	Name     string `json:"username" description:"登录的用户名"`
	Password string `json:"password" description:"登录的密码"`
}

type User struct {
	Id        uint64           `json:"id" description:"用户ID"`
	Ulr       UserLoginRequest `gorm:"embedded"`
	Data      string           `json:"data" description:"用户数据"`
	CreatedAt time.Time        `json:"created_at" description:"更新时间"`
	UpdatedAt time.Time        `json:"updated_at" description:"更新时间"`
}

// func (u *User) Update() error {
// 	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
// 	if err != nil {
// 		return err
// 	}

// 	ret := db.Model(u).Update("name", u.LoginRequest.UserName)
// 	if ret.Error != nil {
// 		return ret.Error
// 	}

// 	return nil
// }

// func (u *User) Delete() error {
// 	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
// 	if err != nil {
// 		return err
// 	}

// 	ret := db.Unscoped().Where("id = ?", u.Id).Delete(u)
// 	if ret.Error != nil {
// 		return ret.Error
// 	}

// 	return nil
// }

func (u *User) BeforeDelete(tx *gorm.DB) error {
	ur := UserRole{
		UserId: u.Id,
	}

	err := ur.Delete()
	if err != nil {
		return err
	}
	return nil
}
