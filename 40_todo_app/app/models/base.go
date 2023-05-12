package models

import (
	"database/sql"
	"fmt"
	"golang_udemy1/40_todo_app/config"
	"log"

	_ "github.com/mattn/go-sql/sqlite3"
)

//テーブル作成

//Userモデルを作成

var Db *sql.DB

var err error

const (
	tableNameUser = "users"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	//Userテーブル作成SQLコマンド
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)

}