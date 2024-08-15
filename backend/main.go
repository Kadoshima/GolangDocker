package main

import (
	"backend/adapter/repository"
	"backend/infrastructure/database"
	"backend/infrastructure/router"
	"backend/migrations"
	"backend/usecase"
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

	// リポジトリの初期化
	userRepo := repository.NewUserRepository(db)

	// ユースケースの初期化
	userUseCase := usecase.NewUserUseCase(userRepo)

	// ルートの設定（依存性注入）
	r := router.SetupRouter(db, userUseCase)
	log.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
