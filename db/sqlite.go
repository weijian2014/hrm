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

	// 创建表
	sql := fmt.Sprintf(`
	CREATE TABLE '%v' (
		'id' INTEGER PRIMARY KEY AUTOINCREMENT,
		'username' TEXT NOT NULL,
		'password' TEXT NOT NULL,
		'type' INTEGER NOT NULL,
		'created' TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		'modify' TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`, LoginTableName)

	_, err = db.Exec(sql)
	if err != nil {
		return err
	}

	log.System("Table %v was created", LoginTableName)

	//插入数据
	sql = fmt.Sprintf(`
	INSERT INTO %v(username, password, type) values(?,?,?)
	`, LoginTableName)
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec("admin", "12345678", 0)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	log.System("Insert admin to teble %v ok, last insert id %v", LoginTableName, id)

	return nil
}
