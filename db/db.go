package db

import (
	"database/sql"
	"errors"
	"fmt"
	"hrm/common"
)

var (
	Conn *sql.DB
)

func openPostgres() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		common.JsonConfigs.PostgresqlHost,
		common.JsonConfigs.PostgresqlPort,
		common.JsonConfigs.PostgresqlUser,
		common.JsonConfigs.PostgresqlPassword,
		common.JsonConfigs.PostgresqlDatabaseName)

	db, err := sql.Open("postgres", psqlInfo)
	if nil != err {
		return nil, err
	}
	// defer db.Close()
	err = db.Ping()
	if nil != err {
		return nil, err
	}

	return db, nil
}

func openSqlite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", common.JsonConfigs.DatabaseName)
	if nil != err {
		return nil, err
	}
	// defer db.Close()
	err = db.Ping()
	if nil != err {
		return nil, err
	}

	return db, nil
}

func openMysql() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		common.JsonConfigs.PostgresqlHost,
		common.JsonConfigs.PostgresqlPort,
		common.JsonConfigs.PostgresqlUser,
		common.JsonConfigs.PostgresqlPassword,
		common.JsonConfigs.PostgresqlDatabaseName)
	db, err := sql.Open("mysql", psqlInfo)
	if nil != err {
		return nil, err
	}
	// defer db.Close()
	err = db.Ping()
	if nil != err {
		return nil, err
	}

	return db, nil
}

// conn需要用户自己关闭
func Open(dbType int) (*sql.DB, error) {
	if nil != Conn {
		return Conn, nil
	}

	var err error
	switch dbType {
	case common.PostgresDatabaseType:
		Conn, err = openPostgres()
	case common.SqliteDatabaseType:
		Conn, err = openSqlite()
	case common.MysqlDatabaseType:
		Conn, err = openMysql()
	}

	if nil != err {
		return nil, errors.New("connect to postgresql fail")
	}

	return Conn, nil
}

func Close() error {
	if nil == Conn {
		return nil
	}

	return Conn.Close()
}
