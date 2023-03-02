package db

import (
	"database/sql"
	"hrm/common"
)

var (
	LoginTableName = "login"
)

// db需要用户自己关闭
func Open(dbType int) (*sql.DB, error) {
	var err error
	var db *sql.DB
	switch dbType {
	case common.PostgresDatabaseType:
		db, err = openPostgres()
	case common.SqliteDatabaseType:
		db, err = openSqlite()
	case common.MysqlDatabaseType:
		db, err = openMysql()
	}

	if nil != err {
		return nil, err
	}

	return db, nil
}

func Init(adminUsername, adminPassword string) error {
	err := createLoginTable()
	if err != nil {
		return err
	}

	err = insertIntoLogin(adminUsername, adminPassword)
	if err != nil {
		return err
	}

	return nil
}
