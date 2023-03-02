package db

import (
	"database/sql"
	"fmt"
	"hrm/common"
	"hrm/log"

	_ "github.com/mattn/go-sqlite3"
)

func openSqlite() (*sql.DB, error) {
	dbPath := common.CurrDir + "/" + common.JsonConfigs.DatabaseName
	// 如果数据库不存在, 则会创建
	db, err := sql.Open("sqlite3", dbPath)
	if nil != err {
		return nil, fmt.Errorf("can not open database %v, %v", dbPath, err.Error())
	}

	// 由用户关闭
	// defer db.Close()
	err = db.Ping()
	if nil != err {
		return nil, err
	}

	return db, nil
}

func createUserTable() error {
	db, err := openSqlite()
	if err != nil {
		return err
	}
	defer db.Close()

	createUserTableStr := `
	CREATE TABLE 'login' (
		'id' INTEGER PRIMARY KEY AUTOINCREMENT,
		'username' TEXT NOT NULL,
		'password' TEXT NOT NULL,
		'type' INTEGER NOT NULL,
		'created' TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		'modify' TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err = db.Exec(createUserTableStr)
	if err != nil {
		return err
	}

	log.System("Table 'login' was created")
	return nil
}
