package api

import (
	"testing"
)

const (
	testExcelFile = "/root/weijian/hrm/doc/template.xlsx"
)

func init() {

}

func Test_ReadExcel(t *testing.T) {
	t.Logf("Test excel file reading function")

	ei, err := ReadExcel(testExcelFile)
	if nil != err {
		t.Error(err)
	}

	t.Logf("Read excel file [%v] ok, content [%v]", testExcelFile, ei)
}

func Test_WriteExcel(t *testing.T) {
	t.Logf("Test excel file writing function")

	ei := &EmployeeInfo{
		Name:            "李四",
		Gender:          "女",
		Age:             "25",
		WorkTime:        "3年",
		Salary:          "3800",
		Post:            "开发部",
		SocialSecurity:  "无",
		Phone:           "13866666666",
		FormerEmployer:  "蓝盾",
		Height:          "178",
		Weight:          "72",
		Diploma:         "高职",
		PoliticalStatus: "党员",
		Id:              "7654321",
		SecurityCard:    "54321",
		CurrentAddress:  "广西省南宁市江南区华夏街道180号",
		Comments:        "1、近日,全国两会受到了社会各界的广泛关注",
	}

	err := WriteExcel(testExcelFile, ei)
	if nil != err {
		t.Error(err)
	}

	t.Logf("Write excel file [%v] ok, content [%v]", testExcelFile, ei)
}
