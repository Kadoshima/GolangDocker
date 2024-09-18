package seeder

import (
	"backend/infrastructure/database"
	"log"
	"os"
)

func Seed() {
	// データベースに接続
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// シード用SQLファイルの読み込み
	seedSQL, err := os.ReadFile("/user.sql")
	if err != nil {
		log.Fatalf("Could not read seed file: %v", err)
	}

	// シードの実行
	if _, err := db.Exec(string(seedSQL)); err != nil {
		log.Fatalf("Could not execute seed SQL: %v", err)
	}

	log.Println("Seed data inserted successfully!")
}

func main() {
	Seed()
}
