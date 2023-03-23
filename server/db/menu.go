package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (m *Menu) CreateTable() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	// 根据Menu结构体，自动创建表结构, 表名为menus
	if err = db.AutoMigrate(m); err != nil {
		return err
	}

	return nil
}

func (m *Menu) Find() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.First(m, "name = ?", m.Name)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (m *Menu) Insert() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Create(m)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (m *Menu) Update() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Model(m).Update("name", m.Name)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (m *Menu) Delete() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Model(m).Delete("name", m.Name)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}
