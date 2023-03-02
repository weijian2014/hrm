package api

import (
	"errors"
	"fmt"
	"hrm/common"
	"hrm/db"
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

	sql := fmt.Sprintf("SELECT password, user_type FROM t_user WHERE user_name='%s';", info.UserName)

	row, err := conn.Query(sql)
	if nil != err {
		return false, err
	}
	defer row.Close()

	if !row.Next() {
		return false, errors.New("用户名或密码错误!")
	}

	dbPassword := ""
	userType := 1
	row.Scan(&dbPassword, &userType)
	if info.Password != dbPassword {
		return false, errors.New("密码错误!")
	}

	// 用户类型 0:管理员 1普通用户
	if userType == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
