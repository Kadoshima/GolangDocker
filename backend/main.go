package main

import (
	"backend/infrastructure/database"
	"backend/infrastructure/router"
	"backend/migrations"
	"log"
	"net/http"
)

func main() {

	// DB初期化
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("Could not connect = %v", err)
	}
	defer db.Close()

	// DBマイグレーション
	migrations.Migrate()

	// ルートの設定
	r := router.SetupRouter()
	log.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
