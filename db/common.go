package db

import (
	"database/sql"
	"fmt"
	"hrm/common"
	"os"
)

var (
	DatabaseName   = "hrm.db"
	LoginTableName = "login"
)

func Init(adminUsername, adminPassword string) error {
	// 数据库文件存在时, 将删除
	dbPath := common.CurrDir + "/" + DatabaseName
	_, err := os.Stat(dbPath)
	if err == nil {
		// 文件存在
		err = os.Remove(dbPath)
		if err != nil {
			return err
		}
		fmt.Printf("Database [%v] has been removed\n", dbPath)
	}

	err = createLoginTable()
	if err != nil {
		return err
	}

	err = insertIntoLogin(adminUsername, adminPassword)
	if err != nil {
		return err
	}

	return nil
}

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
	default:
		panic("Unknown database type")
	}

	if nil != err {
		return nil, err
	}

	return db, nil
}
