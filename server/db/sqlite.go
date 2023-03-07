package db

import (
	"database/sql"
	"fmt"
	"hrm/common"

	_ "github.com/mattn/go-sqlite3"
)

func openSqlite() (*sql.DB, error) {
	dbPath := common.CurrDir + "/" + DatabaseName
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

func createLoginTable() error {
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

	fmt.Printf("Table [%v] was created\n", LoginTableName)
	return nil
}

func insertIntoLogin(adminUsername, adminPassword string) error {
	db, err := openSqlite()
	if err != nil {
		return err
	}
	defer db.Close()

	//插入数据
	sql := fmt.Sprintf(`INSERT INTO %v(username, password, type) values(?,?,?)`, LoginTableName)
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(adminUsername, adminPassword, 0)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	fmt.Printf("Insert admin account to teble [%v] ok, username [%v], password [%v], entry id [%v]\n", LoginTableName, adminUsername, adminPassword, id)
	return nil
}
