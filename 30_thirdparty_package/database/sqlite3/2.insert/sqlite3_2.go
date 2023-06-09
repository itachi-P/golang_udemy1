package main

import (
	"database/sql"
	"log"

	//コードの中では使用しないが必要なインポート
	_ "github.com/mattn/go-sqlite3"
)

//database + sqlite3
//CRUD処理

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	//データの追加
	//VALUES()の"?"がSQLインジェクション(悪意のあるコマンド)をエスケープしてくれる
	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	//VALUESの中の"?"を置換して実際にデータを挿入する
	_, err := Db.Exec(cmd, "世良田次郎三郎元信", 45)
	if err != nil {
		log.Fatalln(err)
	}
}
