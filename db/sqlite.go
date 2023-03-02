package db

import (
	"database/sql"
	"hrm/common"

	"github.com/mattn/go-sqlite3"
)

func openSqlite() (*sql.DB, error) {
	sqlite3.SQLiteDriver

	db, err := sql.Open("sqlite3", common.JsonConfigs.DatabaseName)
	if nil != err {
		return nil, err
	}

	// 由用户关闭
	// defer db.Close()
	err = db.Ping()
	if nil != err {
		return nil, err
	}

	return db, nil
}
