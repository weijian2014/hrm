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

	users := []User{
		{
			Name:     adminUsername,
			Password: adminPassword,
			Data:     "json data",
		},
		{
			Name:     "test",
			Password: "123456",
			Data:     "json data",
		},
	}
	err = users[0].CreateTable()
	if err != nil {
		return err
	}
	fmt.Printf("Table name [users] has been created\n")
	for _, u := range users {
		err = u.Insert()
		if err != nil {
			return err
		}
		fmt.Printf("Inserted row[%v] into table [users]\n", u)
	}

	roles := []Role{
		{
			Name: "超级管理员",
		},
		{
			Name: "普通员工",
		},
	}
	err = roles[0].CreateTable()
	if err != nil {
		return err
	}
	fmt.Printf("Table name [roles] has been created\n")
	for _, r := range roles {
		err = r.Insert()
		if err != nil {
			return err
		}
		fmt.Printf("Inserted row[%v] into table [roles]\n", r)
	}

	menus := []Menu{
		{
			Name: "员工管理",
			Url:  "url",
		},
		{
			Name: "系统管理",
			Url:  "url",
		},
	}
	err = menus[0].CreateTable()
	if err != nil {
		return err
	}
	fmt.Printf("Table name [menus] has been created\n")
	for _, m := range menus {
		err = m.Insert()
		if err != nil {
			return err
		}
		fmt.Printf("Inserted row[%v] into table [menus]\n", m)
	}

	ur := []UserRole{
		{
			UserId: 1,
			RoleId: 1,
		},
		{
			UserId: 2,
			RoleId: 2,
		},
	}
	err = ur[0].CreateTable()
	if err != nil {
		return err
	}
	fmt.Printf("Table name [userroles] has been created\n")
	for _, item := range ur {
		err = item.Insert()
		if err != nil {
			return err
		}
		fmt.Printf("Inserted row[%v] into table [userroles]\n", item)
	}

	rm := []RoleMenu{
		{
			RoleId:   1,
			MenuId:   1,
			ParentId: 0,
		},
		{
			RoleId:   1,
			MenuId:   2,
			ParentId: 0,
		},
		{
			RoleId:   2,
			MenuId:   1,
			ParentId: 0,
		},
	}
	err = rm[0].CreateTable()
	if err != nil {
		return err
	}
	fmt.Printf("Table name [role_menus] has been created\n")
	for _, item := range rm {
		err = item.Insert()
		if err != nil {
			return err
		}
		fmt.Printf("Inserted row[%v] into table [role_menus]\n", item)
	}

	return nil
}
