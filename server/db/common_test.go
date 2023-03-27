package db

import (
	"os"
	"testing"
)

func init() {

}

type Test struct {
	Id   uint
	Name string
	Age  uint
}

func Test_Create_Insert(t *testing.T) {
	test := &Test{Name: "hrm", Age: 18}

	// 数据库文件存在时, 将删除
	_, err := os.Stat(DatabaseFullPath)
	if err == nil {
		// 文件存在
		err = os.Remove(DatabaseFullPath)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}

	// CreateTable
	err = CreateTable(test)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Create table ok")

	err = Find(test, -1, "id = ?", test.Id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// Insert
	err = Insert(test)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Insert ok, [%v]", test)

	// UpdateSingleColumn
	err = UpdateSingleColumn(test, "age", 20)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if test.Age != 20 {
		t.Errorf("No expected value")
		t.FailNow()
	}

	test2 := &Test{Name: "hrm2"}
	err = Take(test2, "id = ?", test.Id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if test2.Age != 20 {
		t.Errorf("No expected value")
		t.FailNow()
	}
	t.Logf("Update single column ok, [%v]", test)

	// UpdateRow
	err = UpdateRow(test)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	test2.Age = 99
	test2.Name = "unknown"
	err = Take(test2, "id = ?", test.Id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if test2.Age != 20 {
		t.Errorf("No expected value")
		t.FailNow()
	}
	t.Logf("Update row ok, [%v]", test)

	// SeleteAll
	test3 := &Test{Name: "hrm3", Age: 22}
	err = Insert(test3)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	tests := new([]Test)
	err = SelectAll(tests)
	if err != nil || len(*tests) != 2 {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("SeleteAll ok, [%v]", test)

	// Delete
	err = Delete(test, "id = ?", test.Id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	err = Find(test2, -1, "id = ?", test.Id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Delete row ok, [%v]", test)

	t.Logf("Test ok")
}
