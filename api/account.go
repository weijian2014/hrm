package api

import (
	"errors"
	"fmt"
	"hrm/common"
	"hrm/db"
)

var (
	loginTableName = "login"
)

type LoginInfo struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func Login(info *LoginInfo) (bool, error) {
	conn, err := db.Open(common.DbType)
	if nil != err {
		return false, err
	}
	defer conn.Close()

	sql := fmt.Sprintf("SELECT password, type FROM %v WHERE username='%v';", loginTableName, info.UserName)
	rows, err := conn.Query(sql)
	if err != nil || !rows.Next() {
		return false, errors.New("用户名或密码错误")
	}

	password := ""
	userType := 1
	rows.Scan(&password, &userType)
	if info.Password != password {
		return false, errors.New("密码错误")
	}

	// 用户类型 0:管理员 1普通用户
	if userType == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
