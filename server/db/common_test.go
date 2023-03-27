package db

import (
	"testing"
)

func init() {

}

func Test_Update(t *testing.T) {
	user := User{
		Id: 2,
		Ulr: UserLoginRequest{
			Name:     "test",
			Password: "123456678",
		},
	}

	// map[string]interface{}{"name": "test", "password": "12345678"},
	cols := []string{"password"}
	err := UpdateColumns(user, cols)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("ok, [%v]", user)
}
