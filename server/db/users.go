package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	Name     string
	Password string
	Data     string
	gorm.Model
}

func FindUser(key interface{}) (*User, error) {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	u := new(User)
	r := db.First(u, "name = ?", key)
	if r.Error != nil {
		return nil, r.Error
	}

	return u, nil
}
