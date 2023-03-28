package db

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"gorm.io/gorm"
)

var (
	user1 = &User{Name: "user1", Password: "123456", Data: "user_data1"}
	user2 = &User{Name: "user2", Password: "123456", Data: "user_data2"}
	user3 = &User{Name: "user1", Password: "123456", Data: "user_data1"}
)

func setup() {
	fmt.Println("Before all tests")

	// 数据库文件存在时, 将删除
	_, err := os.Stat(DatabaseFullPath)
	if err == nil {
		// 文件存在
		err = os.Remove(DatabaseFullPath)
		if err != nil {
			panic(err)
		}
	}

	user := &User{}
	err = CreateTable(user)
	if err != nil {
		panic(err)
	}

	err = Insert(&user1)
	if err != nil {
		panic(err)
	}

	err = Insert(&user2)
	if err != nil {
		panic(err)
	}

	err = Insert(&user3)
	if err != nil {
		panic(err)
	}
}

func teardown() {
	fmt.Println("After all tests")

	// 数据库文件存在时, 将删除
	_, err := os.Stat(DatabaseFullPath)
	if err == nil {
		// 文件存在
		err = os.Remove(DatabaseFullPath)
		if err != nil {
			panic(err)
		}
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func Test_CreateTable(t *testing.T) {
	// 创建单个表
	role := &Role{}
	err := CreateTable(role)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// 创建多个表
	menu := &Menu{}
	employee := &Employee{}
	err = CreateTable(menu, employee)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Create table [roles, menus, employees] ok")

	// 重复创建, 先删除原来的表
	err = CreateTable(role)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("Test ok")
}

func Test_Insert(t *testing.T) {
	// 插入单行
	test1 := &User{Name: "test1", Password: "123456", Data: "test_data1"}

	err := Insert(test1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Insert row[%+v] ok", test1)

	// 插入数组
	tests := []User{
		{Name: "test2", Password: "123456", Data: "test_data2"},
		{Name: "test3", Password: "123456", Data: "test_data3"},
	}
	err = Insert(&tests)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Insert rows[%+v] ok", tests)

	// id已经存在, 插入失败
	err = Insert(&tests)
	if err == nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("Test ok")
}

func Test_Fisrt(t *testing.T) {
	output1 := new(User)
	err := First(output1, "name = ?", user1.Name)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output1.Id != user1.Id || output1.Name != user1.Name || output1.Data != user1.Data || output1.Password != user1.Password {
		t.Errorf("First error")
		t.FailNow()
	}
	t.Logf("First row[%+v]", output1)

	output2 := new(User)
	err = First(output2, "name = ?", user2.Name)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output2.Id != user2.Id || output2.Name != user2.Name || output2.Data != user2.Data || output2.Password != user2.Password {
		t.Errorf("First error")
		t.FailNow()
	}
	t.Logf("First row[%+v]", output2)

	// 只使用id(主键), 不需要其它查询条件即可找到
	output3 := &User{Id: 3}
	err = First(output3, "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("First row[%+v]", output3)

	// 找不到记录, 返回 ErrRecordNotFound 错误
	output4 := new(User)
	err = First(output4, "name = ?", "unknown")
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("Test ok")
}

func Test_Last(t *testing.T) {
	output1 := new(User)
	// 最后匹配的行
	err := Last(output1, "name = ?", user1.Name)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output1.Id != user3.Id || output1.Name != user3.Name || output1.Data != user3.Data || output1.Password != user3.Password {
		t.Errorf("Last error")
		t.FailNow()
	}
	t.Logf("Last row[%+v]", output1)

	output2 := new(User)
	err = Last(output2, "name = ?", user2.Name)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output2.Id != user2.Id || output2.Name != user2.Name || output2.Data != user2.Data || output2.Password != user2.Password {
		t.Errorf("Last error")
		t.FailNow()
	}
	t.Logf("Last row[%+v]", output2)

	// 只使用id(主键), 不需要其它查询条件即可找到
	output3 := &User{Id: 1}
	err = Last(output3, "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Last row[%+v]", output3)

	// 找不到记录, 返回 ErrRecordNotFound 错误
	output4 := new(User)
	err = Last(output4, "name = ?", "unknown")
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("Test ok")
}

func Test_Take(t *testing.T) {
	output1 := new(User)
	// 不排序, 第一条匹配的行
	err := Take(output1, "name = ?", user1.Name)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output1.Id != user1.Id || output1.Name != user1.Name || output1.Data != user1.Data || output1.Password != user1.Password {
		t.Errorf("Take error")
		t.FailNow()
	}
	t.Logf("Take row[%+v]", output1)

	output2 := new(User)
	err = Take(output2, "name = ?", user2.Name)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output2.Id != user2.Id || output2.Name != user2.Name || output2.Data != user2.Data || output2.Password != user2.Password {
		t.Errorf("Take error")
		t.FailNow()
	}
	t.Logf("Take row[%+v]", output2)

	// 只使用id(主键), 不需要其它查询条件即可找到
	output3 := &User{Id: 1}
	err = Take(output3, "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Take row[%+v]", output3)

	// 找不到记录, 返回 ErrRecordNotFound 错误
	output4 := new(User)
	err = Last(output4, "name = ?", "unknown")
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("Test ok")
}

func Test_Find(t *testing.T) {
	output1 := new(User)
	// 第一条匹配的行
	err := Find(output1, 1, "name = ?", user1.Name)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output1.Id != user1.Id || output1.Name != user1.Name || output1.Data != user1.Data || output1.Password != user1.Password {
		t.Errorf("Find error")
		t.FailNow()
	}
	t.Logf("Find row[%+v]", output1)

	output2 := new(User)
	err = Find(output2, -1, "name = ?", user2.Name)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output2.Id != user2.Id || output2.Name != user2.Name || output2.Data != user2.Data || output2.Password != user2.Password {
		t.Errorf("Find error")
		t.FailNow()
	}
	t.Logf("Find row[%+v]", output2)

	// 只使用id(主键), 不需要其它查询条件即可找到
	output3 := &User{Id: 1}
	err = Find(output3, 1, "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Find row[%+v]", output3)

	// 找不到记录, 返回 ErrRecordNotFound 错误
	output4 := new(User)
	err = Find(output4, 1, "name = ?", "unknown")
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error(err)
		t.FailNow()
	}

	// 查找多条1
	outputs1 := new([]User)
	err = Find(&outputs1, 3, "name = ?", user1.Name)
	if err != nil || len(*outputs1) != 2 {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Find rows[%+v]", outputs1)

	outputs2 := new([]User)
	err = Find(&outputs2, -1)
	if err != nil || len(*outputs2) != 3 {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Find rows[%+v]", outputs2)

	t.Logf("Test ok")
}
