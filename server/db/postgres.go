package db

import (
	"database/sql"
	"fmt"
	"hrm/common"
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

	// 由用户关闭
	// defer db.Close()
	err = db.Ping()
	if nil != err {
		return nil, err
	}

	return db, nil
}
