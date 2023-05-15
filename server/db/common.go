package db

import (
	"encoding/json"
	"fmt"
	"hrm/common"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DatabaseName     = "hrm.db"
	DatabaseFullPath = common.CurrDir + "/" + DatabaseName
)

type Gen struct {
	Name             string
	Phone            string
	Identifier       string
	Address          string
	Company          string
	Email            string
	CreditCardNumber string `json:"credit_card_number"`
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

	// 表users
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
	err = CreateTable(&User{})
	if err != nil {
		return err
	}
	fmt.Printf("Table [users] has been created\n")
	err = Insert(&users)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted rows into table [users]\n")

	// 表roles
	roles := []Role{
		{
			Name: "超级管理员",
		},
		{
			Name: "普通员工",
		},
	}
	err = CreateTable(&Role{})
	if err != nil {
		return err
	}
	fmt.Printf("Table [roles] has been created\n")
	err = Insert(&roles)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted rows into table [roles]\n")

	// 表menus
	menus := []Menu{
		{
			Id:       1,
			Name:     "员工管理",
			ParentId: 0,
			Icon:     "lock",
			Url:      "url",
		},
		{
			Id:       2,
			Name:     "招聘管理",
			ParentId: 0,
			Icon:     "lock",
			Url:      "url",
		},
		{
			Id:       3,
			Name:     "权限管理",
			ParentId: 0,
			Icon:     "lock",
			Url:      "url",
		},
		{
			Id:       4,
			Name:     "系统管理",
			ParentId: 0,
			Icon:     "lock",
			Url:      "url",
		},
		{
			Id:       5,
			Name:     "菜单列表",
			ParentId: 3,
			Icon:     "lock",
			Url:      "url",
		},
		{
			Id:       6,
			Name:     "角色列表",
			ParentId: 3,
			Icon:     "lock",
			Url:      "url",
		},
		{
			Id:       7,
			Name:     "用户列表",
			ParentId: 3,
			Icon:     "lock",
			Url:      "url",
		},
	}
	err = CreateTable(&Menu{})
	if err != nil {
		return err
	}
	fmt.Printf("Table [menus] has been created\n")
	err = Insert(&menus)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted rows into table [menus]\n")

	// 表user_roles
	urs := []UserRole{
		{
			UserId: 1,
			RoleId: 1,
		},
		{
			UserId: 2,
			RoleId: 2,
		},
	}
	err = CreateTable(&UserRole{})
	if err != nil {
		return err
	}
	fmt.Printf("Table [user_roles] has been created\n")
	err = Insert(&urs)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted rows into table [user_roles]\n")

	// 表role_menus
	rms := []RoleMenu{
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
	err = CreateTable(&RoleMenu{})
	if err != nil {
		return err
	}
	fmt.Printf("Table [role_menus] has been created\n")
	err = Insert(&rms)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted rows into table [role_menus]\n")

	// 表posts
	p := []Post{
		{
			Name: "前台",
		},
		{
			Name: "巡逻",
		},
		{
			Name: "消防",
		},
		{
			Name: "教练",
		},
		{
			Name: "保安",
		},
		{
			Name: "银保",
		},
	}
	err = CreateTable(&Post{})
	if err != nil {
		return err
	}
	fmt.Printf("Table [posts] has been created\n")
	err = Insert(&p)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted rows into table [posts]\n")

	// 表employees
	const count = 100
	var employees [count]Employee

	jsonData, err := common.Command(common.CurrDir+"/doc/faker_gen.py", strconv.Itoa(count))
	if err != nil {
		return err
	}
	fmt.Printf("jsonData=[%v]\n", string(jsonData))

	var g []Gen
	if err = json.Unmarshal(jsonData, &g); err != nil {
		return err
	}

	for i := range employees {
		employees[i].Name = g[i].Name
		if i%2 == 0 {
			employees[i].Gender = "男"
		} else {
			employees[i].Gender = "女"
		}

		// 20~55岁
		age := rand.Intn(55-20) + 20
		employees[i].Birthday = time.Now().AddDate(-age, 0, 0)
		employees[i].FirstWorkTime = time.Now().AddDate(-(age - 20), 0, 0)
		employees[i].Salary = uint64(age * 200)
		employees[i].PostId = uint64(rand.Intn(len(p)-1) + 1)
		ss := []string{"有", "无"}
		employees[i].SocialSecurity = ss[rand.Intn(len(ss))]
		employees[i].Phone = fmt.Sprintf("138%v%v", rand.Intn(9999-1000)+1000, rand.Intn(9999-1000)+1000)
		employees[i].FormerEmployer = fmt.Sprintf("原单位%v-%v", i, g[i].Email)
		employees[i].Height = uint64(rand.Intn(220-150) + 150)
		employees[i].Weight = uint64(rand.Intn(120-45) + 45)
		d := []string{"小学", "初中", "高中", "中专", "大专", "专升本", "本科", "研究生", "博士", "博士后"}
		employees[i].Degree = d[rand.Intn(len(d))]
		ps := []string{"群众", "无党派人士", "民主党派成员", "共青团员", "党员"}
		employees[i].PoliticalStatus = ps[rand.Intn(len(ps))]
		employees[i].Identifier = fmt.Sprintf("%v%v%v", g[i].Identifier[0:6], employees[i].Birthday.Format("20060102"), g[i].Identifier[14:])
		employees[i].SecurityCard = strconv.Itoa(rand.Intn(99999999-10000000) + 10000000)
		employees[i].CurrentAddress = g[i].Address
		employees[i].Comments = g[i].CreditCardNumber + " " + g[i].Email
		// fmt.Printf("employee=[%v]\n", employees[i])
	}
	err = CreateTable(&Employee{})
	if err != nil {
		return err
	}
	fmt.Printf("Table [employees] has been created\n")
	err = Insert(&employees)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted rows[%+v] into table [employees]\n", len(employees))

	return nil
}

// 表存在时先删除原来的表
func CreateTable(dst ...interface{}) error {
	db, err := getDb()
	if err != nil {
		return err
	}

	return db.AutoMigrate(dst...)
}

// 获取表中满足conds的主键升序的第一条记录
// 找不到记录时返回 ErrRecordNotFound 错误
func First(output interface{}, query interface{}, args ...interface{}) error {
	db, err := getDb()
	if err != nil {
		return err
	}

	result := db.Where(query, args...).First(output)
	return result.Error
}

// 获取表中满足conds的主键降序的最后一条记录
// 找不到记录时返回 ErrRecordNotFound 错误
func Last(output interface{}, query interface{}, args ...interface{}) error {
	db, err := getDb()
	if err != nil {
		return err
	}

	result := db.Where(query, args...).Last(output)
	return result.Error
}

// 获取没有指定排序字段的第一条记录
// 找不到记录时返回 ErrRecordNotFound 错误
func Take(output interface{}, query interface{}, args ...interface{}) error {
	db, err := getDb()
	if err != nil {
		return err
	}

	result := db.Where(query, args...).Take(output)
	return result.Error
}

// 根据conds查找
// 找不到记录时不返回错误
// limit 为限制获取的行数, -1表示不限制
func Find1(output interface{}, limit int, conds ...interface{}) error {
	db, err := getDb()
	if err != nil {
		return err
	}

	result := db.Limit(limit).Find(output, conds...)
	if limit != -1 && result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	} else {
		return result.Error
	}
}

// 根据conds查找
// 找不到记录时不返回错误
// join 连接的表名 string
// joinConds 连接的条件 string
// limit 为限制获取的行数, -1表示不限制
// 查询条件
func Find2(output interface{}, join string, joinConds interface{}, limit int, conds ...interface{}) error {
	db, err := getDb()
	if err != nil {
		return err
	}

	result := db.Joins(join, joinConds).Limit(limit).Find(output, conds...)
	if limit != -1 && result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	} else {
		return result.Error
	}
}

// output 可以是数组, 实现批量插入
// 插入成功后output会返回行的主键值
func Insert(output interface{}) error {
	db, err := getDb()
	if err != nil {
		return err
	}

	result := db.Create(output)
	return result.Error
}

// 更新整行所有列
// Save() 不要和 Model()一起使用
func UpdateRow(dst interface{}) error {
	db, err := getDb()
	if err != nil {
		return err
	}

	result := db.Save(dst)
	return result.Error
}

// 更新单个列
func UpdateSingleColumn(dst interface{}, column string, value interface{}) error {
	db, err := getDb()
	if err != nil {
		return err
	}

	result := db.Model(dst).Update(column, value)
	return result.Error
}

// 更新指定列
// columnsUpdate可以是[]string
func UpdateColumns(dst interface{}, columnsUpdate interface{}) error {
	db, err := getDb()
	if err != nil {
		return err
	}

	result := db.Model(dst).Select(columnsUpdate).Updates(dst)
	return result.Error
}

// 永久删除(非软删除)满足条件的行
func Delete(dst interface{}, conds ...interface{}) error {
	db, err := getDb()
	if err != nil {
		return err
	}

	// Unscoped() 永久删除
	result := db.Unscoped().Delete(dst, conds...)
	return result.Error
}

func getDb() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(DatabaseFullPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
