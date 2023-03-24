package db

import (
	"hrm/log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (ur *UserRole) CreateTable() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	// 表名为user_roles
	if err = db.AutoMigrate(ur); err != nil {
		return err
	}

	return nil
}

func (ur *UserRole) Find() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.First(ur, "user_id = ?", ur.UserId)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (ur *UserRole) Insert() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Create(ur)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (ur *UserRole) Update() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Model(ur).Update("user_id", ur.UserId)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (ur *UserRole) Delete() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Unscoped().Where("user_id = ? or role_id = ?", ur.UserId, ur.RoleId).Delete(ur)
	if ret.Error != nil {
		return ret.Error
	}

	log.Debug("Delete userId[%v] or roleId[%v] form table user_roles", ur.UserId, ur.RoleId)
	return nil
}
