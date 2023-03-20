package db

import (
	"fmt"
	"hrm/common"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DatabaseName = "hrm.db"
)

type User struct {
	Name     string
	Password string
	Data     string
	gorm.Model
}

func Init(adminUsername, adminPassword string) error {
	// 数据库文件存在时, 将删除
	dbPath := common.CurrDir + "/" + DatabaseName
	_, err := os.Stat(dbPath)
	if err == nil {
		// 文件存在
		err = os.Remove(dbPath)
		if err != nil {
			return err
		}
		fmt.Printf("Database [%v] has been removed\n", dbPath)
	}

	db, err := gorm.Open(sqlite.Open(DatabaseName), &gorm.Config{})
	if err != nil {
		return err
	}

	// 根据User结构体，自动创建表结构.
	db.AutoMigrate(&User{})

	// 插入记录
	r := db.Create(&User{Name: adminUsername, Password: adminPassword})
	if r.Error != nil {
		return r.Error
	}

	return nil
}
