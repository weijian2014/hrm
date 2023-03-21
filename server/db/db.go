package db

import (
	"fmt"
	"hrm/common"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DatabaseName     = "hrm.db"
	DatabaseFullPath = common.CurrDir + "/" + DatabaseName
)

func Init(adminUsername, adminPassword string) error {
	// 数据库文件存在时, 将删除
	_, err := os.Stat(DatabaseFullPath)
	if err == nil {
		// 文件存在
		err = os.Remove(DatabaseFullPath)
		if err != nil {
			return err
		}
		fmt.Printf("Database [%v] has been removed\n", DatabaseFullPath)
	}

	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Printf("Database [%v] has been created\n", DatabaseFullPath)

	// 根据User结构体，自动创建表结构, 表名为users
	if err = db.AutoMigrate(&User{}); err != nil {
		return err
	}
	fmt.Printf("Table [users] has been created in [%v]\n", DatabaseFullPath)

	// 插入记录
	if err != nil {
		panic(err)
	}
	r := db.Create(&User{Name: adminUsername, Password: adminPassword, Data: "json data"})
	if r.Error != nil {
		return r.Error
	}
	fmt.Printf("Table [users] insert row[%v, %v, %v]\n", adminUsername, adminPassword, "json data")

	return nil
}
