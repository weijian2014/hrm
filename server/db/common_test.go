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

	// 表存在时先删除原来的表
	err = CreateTable(role)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("Test ok")
}

func Test_Insert(t *testing.T) {
	type TestInsert struct {
		Id   uint32
		Name string
		Data string
	}

	ti := &TestInsert{}
	err := CreateTable(ti)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// 插入单行
	test1 := &TestInsert{Name: "test1", Data: "test_data1"}
	err = Insert(test1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Insert row[%+v] ok", test1)

	// 插入数组
	tests := []TestInsert{
		{Name: "test2", Data: "test_data2"},
		{Name: "test3", Data: "test_data3"},
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
	err := Find1(output1, 1, "name = ?", user1.Name)
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
	err = Find1(output2, -1, "name = ?", user2.Name)
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
	err = Find1(output3, 1, "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Find row[%+v]", output3)

	// 找不到记录, 返回 ErrRecordNotFound 错误
	output4 := new(User)
	err = Find1(output4, 1, "name = ?", "unknown")
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error(err)
		t.FailNow()
	}

	// 查找多条1
	outputs1 := new([]User)
	err = Find1(&outputs1, 3, "name = ?", user1.Name)
	if err != nil || len(*outputs1) != 2 {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Find rows[%+v]", outputs1)

	outputs2 := new([]User)
	err = Find1(&outputs2, -1)
	if err != nil || len(*outputs2) != 3 {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Find rows[%+v]", outputs2)

	t.Logf("Test ok")
}

func Test_Update(t *testing.T) {
	type TestUpdate struct {
		Id   uint32
		Name string
		Age  int
		Data string
	}

	tu := &TestUpdate{}
	err := CreateTable(tu)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	test1 := &TestUpdate{Name: "test1", Age: 18, Data: "test_data1"}
	err = Insert(test1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Insert row[%+v] ok", test1)

	name := "UpdateSingleColumn"
	test2 := &TestUpdate{Id: test1.Id, Name: name}
	err = UpdateSingleColumn(test2, "name", test2.Name)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	output1 := new(TestUpdate)
	err = Take(output1, "id = ?", test1.Id)
	if err != nil || output1.Name != name {
		t.Error(err)
		t.FailNow()
	}

	// 多列更新
	name = "UpdateColumns1"
	data := "UpdateColumnsData1"
	age := 88
	test3 := &TestUpdate{Id: test1.Id, Name: name, Age: age, Data: data}
	err = UpdateColumns(test3, []string{"name", "data"})
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	output2 := new(TestUpdate)
	err = Take(output2, "id = ?", test1.Id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output2.Name != name || output2.Age != test1.Age || output2.Data != data {
		t.Errorf("No expected, [%+v]", output2)
		t.FailNow()
	}

	// 单列更新
	name = "UpdateColumns2"
	data = "UpdateColumnsData2"
	age = 99
	test4 := &TestUpdate{Id: test1.Id, Name: name, Age: age, Data: data}
	err = UpdateColumns(test4, "age")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	output3 := new(TestUpdate)
	err = Take(output3, "id = ?", test1.Id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output3.Name != output2.Name || output3.Data != output2.Data || output3.Age != age {
		t.Errorf("No expected, [%+v]", output3)
		t.FailNow()
	}

	// 整行更新
	name = "UpdateRow"
	data = "UpdateRowData"
	age = 100
	test5 := &TestUpdate{Id: test1.Id, Name: name, Age: age, Data: data}
	err = UpdateRow(test5)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	output4 := new(TestUpdate)
	err = Take(output4, "id = ?", test1.Id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output4.Name != name || output4.Data != data || output4.Age != age {
		t.Errorf("No expected, [%+v]", output4)
		t.FailNow()
	}

	t.Logf("Test ok")
}

func Test_Delete(t *testing.T) {
	type TestDelete struct {
		Id   uint32
		Name string
		Age  int
		Data string
	}

	td := &TestDelete{}
	err := CreateTable(td)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	test1 := &TestDelete{Name: "test1", Age: 11, Data: "test_data1"}
	test2 := &TestDelete{Name: "test2", Age: 22, Data: "test_data2"}
	test3 := &TestDelete{Name: "test3", Age: 33, Data: "test_data3"}
	tests := []TestDelete{*test1, *test2, *test3}
	err = Insert(tests)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Insert rows[%+v] ok", tests)

	// 提供ID(主键), 不需要其它条件
	err = Delete(tests[0])
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	output1 := new(TestDelete)
	err = Take(output1, "id = ?", tests[0].Id)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error(err)
		t.FailNow()
	}

	// 提供条件, 不提供ID(主键)
	err = Delete(&TestDelete{}, "name = ?", test2.Name)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	output2 := new(TestDelete)
	err = Take(output2, "id = ?", tests[1].Id)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error(err)
		t.FailNow()
	}

	//  提供的条件不存在
	err = Delete(&TestDelete{}, "name = ?", "not found")
	if err != nil {
		t.Errorf("%+v", err)
		t.FailNow()
	}
	output3 := new(TestDelete)
	err = Take(output3, "id = ?", tests[2].Id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Take row[%+v] ok", output3)

	t.Logf("Test ok")
}
