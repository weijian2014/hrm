package db

import (
	"testing"
)

func init() {

}

func Test_Like(t *testing.T) {
	employees := new([]Employee)
	err := Find1(employees, -1, "name || height || weight || degree || identifier || phone || current_address || former_employer || salary || security_card || comments like '%?%'", "方")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("ok, [%+v]", employees)
}
