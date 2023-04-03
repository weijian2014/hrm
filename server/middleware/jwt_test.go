package middleware

import (
	"strings"
	"testing"
)

func init() {

}

func Test_GenerateToken_And_IsTokenValid(t *testing.T) {
	t.Logf("Test excel file reading function")
	userId := uint64(1)
	username := "weijian"

	info, err := GenerateToken(userId, username, 60)
	if nil != err {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("Generate user info[%+v]", info)

	ht := strings.Split(info.AccessToken, " ")
	claims, err := isTokenValid(ht[1])
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if claims.UserName != username {
		t.FailNow()
	}

	claims, err = isTokenValid(info.RefreshToken)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if claims.UserName != username {
		t.FailNow()
	}

	t.Logf("Test ok")
}
