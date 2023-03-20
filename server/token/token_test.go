package token

import (
	"testing"
)

func init() {

}

func Test_GenerateToken_And_IsTokenValid(t *testing.T) {
	t.Logf("Test excel file reading function")
	username := "weijian"

	token, err := GenerateToken(username)
	if nil != err {
		t.Error(err)
		t.FailNow()
	}

	t.Logf("Generate token [%v] for username [%v] ok", token, username)

	un, err := IsTokenValid(token)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if un != username {
		t.FailNow()
	}

	t.Logf("Token [%v] is valid for username [%v]", token, un)
}
