package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Password string
	Data     string
}

func (u *User) Find(key interface{}) error {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}

	r := db.First(u, "name = ?", key)
	if r.Error != nil {
		return r.Error
	}

	return nil
}

func (u *User) Update(key interface{}) error {

	return nil
}

func (u *User) Delete(key interface{}) error {

	return nil
}
