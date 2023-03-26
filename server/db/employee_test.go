package db

import (
	"testing"
)

func init() {

}

func Test_Default(t *testing.T) {
	ps := new(PaginationSettings)
	t.Logf("ok, [%v]", ps)
}
