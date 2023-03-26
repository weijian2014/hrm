package db

import (
	"encoding/json"
	"fmt"
	"hrm/common"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DatabaseName     = "hrm.db"
	DatabaseFullPath = common.CurrDir + "/" + DatabaseName

	// test
	api = "http://0.0.0.0:7070/api"
)

type Gen struct {
	Name   string
	Mobile string
	IdNo   string
	Bank   string
	Email  string
	Addr   string
}

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
			Name: "招聘管理",
			Url:  "url",
		},
		{
			Name: "休假管理",
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
			RoleId:   1,
			MenuId:   3,
			ParentId: 0,
		},
		{
			RoleId:   1,
			MenuId:   4,
			ParentId: 0,
		},
		{
			RoleId:   2,
			MenuId:   3,
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

	var employee [100]Employee
	for i := range employee {
		r, err := http.Get(api)
		if err != nil {
			return err
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		// fmt.Printf("data=[%v]\n", string(data))

		var g Gen
		if err = json.Unmarshal(data, &g); err != nil {
			return err
		}

		employee[i].Name = g.Name
		if i%2 == 0 {
			employee[i].Gender = "男"
		} else {
			employee[i].Gender = "女"
		}

		// 20~55岁
		age := rand.Intn(55-20) + 20
		employee[i].Birthday = time.Now().AddDate(-age, 0, 0)
		employee[i].FirstWorkTime = time.Now().AddDate(-(age - 20), 0, 0)
		employee[i].Salary = uint64(age * 200)
		posts := []string{"保安", "消防", "前台", "后勤", "保洁", "经理", "财务", "行政"}
		employee[i].Post = posts[rand.Intn(len(posts))]
		ss := []string{"有", "无"}
		employee[i].SocialSecurity = ss[rand.Intn(len(ss))]
		employee[i].Phone = fmt.Sprintf("138%v%v", rand.Intn(9999-1000)+1000, rand.Intn(9999-1000)+1000)
		employee[i].FormerEmployer = fmt.Sprintf("原单位%v-%v", i, g.Email)
		employee[i].Height = uint64(rand.Intn(220-150) + 150)
		employee[i].Weight = uint64(rand.Intn(120-45) + 45)
		d := []string{"小学", "初中", "高中", "中专", "大专", "专升本", "本科", "研究生", "博士", "博士后"}
		employee[i].Degree = d[rand.Intn(len(d))]
		ps := []string{"群众", "无党派人士", "民主党派成员", "共青团员", "党员"}
		employee[i].PoliticalStatus = ps[rand.Intn(len(ps))]
		employee[i].Identifier = fmt.Sprintf("%v%v%v", g.IdNo[0:6], employee[i].Birthday.Format("20060102"), g.IdNo[14:])
		employee[i].SecurityCard = strconv.Itoa(rand.Intn(99999999-10000000) + 10000000)
		employee[i].CurrentAddress = g.Addr
		employee[i].Comments = g.Bank + " " + g.Email
		// fmt.Printf("employee=[%v]\n", employee[i])
	}

	err = employee[0].CreateTable()
	if err != nil {
		return err
	}
	fmt.Printf("Table name [employees] has been created\n")
	for _, e := range employee {
		err = e.Insert()
		if err != nil {
			return err
		}
		fmt.Printf("Inserted row[%v] into table [employees]\n", e)
	}

	return nil
}
