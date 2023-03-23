package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (r *Role) CreateTable() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	// 根据Role结构体，自动创建表结构, 表名为roles
	if err = db.AutoMigrate(r); err != nil {
		return err
	}

	return nil
}

func (r *Role) Find() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.First(r, "name = ?", r.Name)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (r *Role) Insert() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	// 插入记录
	ret := db.Create(r)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (r *Role) Update() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Model(r).Update("name", r.Name)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (r *Role) Delete() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Model(r).Delete("name", r.Name)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}
