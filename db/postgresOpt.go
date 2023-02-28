package db

import (
	"database/sql"
	"errors"
	"fmt"
	"hrm/common"
)

type PgInfo struct {
	Conn *sql.DB
}

func connect() (*sql.DB, error) {
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

// conn需要用户自己关闭
func New() (*PgInfo, error) {
	db, err := connect()
	if nil != err {
		return nil, errors.New("Connect to postgresql fail")
	}

	c := new(PgInfo)
	c.Conn = db
	return c, nil
}

func (c *PgInfo) Close() error {
	if nil == c || nil == c.Conn {
		return nil
	}

	return c.Conn.Close()
}
