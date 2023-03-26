package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func (u *User) CreateTable() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	// 表名为users
	if err = db.AutoMigrate(u); err != nil {
		return err
	}

	return nil
}

func (u *User) Find() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.First(u, "name = ?", u.Name)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (u *User) Insert() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Create(u)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (u *User) Update() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Model(u).Update("name", u.Name)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

func (u *User) Delete() error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	ret := db.Unscoped().Where("id = ?", u.Id).Delete(u)
	if ret.Error != nil {
		return ret.Error
	}

	return nil
}

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
