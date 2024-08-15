package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	dsn := "root:pass@tcp(db:3306)/test_db" // テスト用なのでgithubに挙げて問題なし
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}

	// データベース接続を確認
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to the database!")
	return db, nil
}
