package db

import (
	"testing"
)

func init() {

}

func Test_Query(t *testing.T) {
	u := &User{
		Name: "test",
	}

	err := u.Delete()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log("ok")
}
