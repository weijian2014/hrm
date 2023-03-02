package session

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_GetMd5(t *testing.T) {
	fmt.Println("Test_GetMd5...")
	h := md5.New()
	h.Write([]byte("weijian.com"))
	fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil)))
	fmt.Println("Test_GetMd5 ok")
}
