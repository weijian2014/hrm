package db

import (
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

	ret := db.Model(ur).Delete("user_id", ur.UserId)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}
