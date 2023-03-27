package db

import (
	"testing"
)

func init() {

}

func Test_Default(t *testing.T) {
	ps := new(EmployeePaginationSetting)
	t.Logf("ok, [%v]", ps)
}
