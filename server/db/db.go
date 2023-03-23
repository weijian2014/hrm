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

	_, err = gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Printf("Database [%v] has been created\n", DatabaseFullPath)

	u := &User{
		Name:     adminUsername,
		Password: adminPassword,
		Data:     "json data",
	}

	err = u.CreateTable()
	if err != nil {
		return err
	}
	fmt.Printf("Table name [users] has been created\n")

	err = u.Insert()
	if err != nil {
		return err
	}
	fmt.Printf("Inserted row[%v] into table [users]\n", u)

	return nil
}
