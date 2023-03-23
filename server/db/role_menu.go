package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (rm *RoleMenu) CreateTable() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	// 表名为role_menus
	if err = db.AutoMigrate(rm); err != nil {
		return err
	}

	return nil
}

func (rm *RoleMenu) Find() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.First(rm, "menu_id = ?", rm.MenuId)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (rm *RoleMenu) Insert() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Create(rm)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (rm *RoleMenu) Update() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Model(rm).Update("menu_id", rm.MenuId)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (rm *RoleMenu) Delete() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Model(rm).Delete("menu_id", rm.MenuId)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}
