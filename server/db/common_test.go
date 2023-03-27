package db

import (
	"os"
	"testing"

	"gorm.io/gorm"
)

func init() {

}

type Test struct {
	Id   uint
	Name string
	Age  uint
}

func Test_Interface(t *testing.T) {
	test1 := &Test{Name: "hrm1", Age: 18}

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
	err = CreateTable(test1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Create table ok")

	output1 := new(Test)
	err = Take(output1, "id = ?", test1.Id)
	if err != gorm.ErrRecordNotFound {
		t.Error(err)
		t.FailNow()
	}

	// Insert
	err = Insert(test1)
	if err != nil || test1.Id != 1 {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Insert ok, [%v]", test1)

	// UpdateSingleColumn
	err = UpdateSingleColumn(test1, "age", 20)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if test1.Age != 20 {
		t.Errorf("No expected value")
		t.FailNow()
	}

	output2 := new(Test)
	err = Take(output2, "id = ?", test1.Id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output2.Age != 20 {
		t.Errorf("No expected value")
		t.FailNow()
	}
	t.Logf("Update single column ok")

	// UpdateRow
	test1.Age = 88
	err = UpdateRow(test1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	output3 := new(Test)
	err = Take(output3, "id = ?", test1.Id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if output3.Age != 88 {
		t.Errorf("No expected value")
		t.FailNow()
	}
	t.Logf("Update row ok")

	// SeleteAll
	test2 := &Test{Name: "hrm3", Age: 33}
	err = Insert(test2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Insert ok, [%v]", test2)
	tests := new([]Test)
	err = SelectAll(tests)
	if err != nil || len(*tests) != 2 {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("SeleteAll ok")

	// Delete
	err = Delete(test1, "id = ?", test1.Id)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	output4 := new(Test)
	err = First(output4, "id = ?", test1.Id)
	if err != gorm.ErrRecordNotFound {
		t.Error(err)
		t.FailNow()
	}

	output5 := new(Test)
	err = Find(output5, 1, "id = ?", test2.Id)
	if err != nil || output5.Age != test2.Age {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("Delete row ok")

	t.Logf("Test ok")
}
