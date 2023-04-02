package middleware

import (
	"testing"
)

func init() {

}

func Test_GenerateToken_And_IsTokenValid(t *testing.T) {
	t.Logf("Test excel file reading function")
	username := "weijian"

	info, err := GenerateToken(username, 60)
	if nil != err {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("Generate user info[%+v]", info)

	claims, err := isTokenValid(info.Token)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if claims.Info.UserName != username {
		t.FailNow()
	}

	t.Logf("Test ok")
}
