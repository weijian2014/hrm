package api

import (
	"fmt"
	"testing"
)

func init() {

}

func Test_ReadExcel(t *testing.T) {
	fmt.Println("Test excel file reading function")

	excelFile := "C:/msp/hrm/doc/template.xlsx"
	ei, err := ReadExcel(excelFile)
	if nil != err {
		panic(err)
	}

	t.Logf("Read excel file [%v]: %v", excelFile, ei)
}
