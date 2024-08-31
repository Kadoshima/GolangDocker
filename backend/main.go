package main

import (
	"backend/adapter/repository"
	"backend/infrastructure/auth"
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

	// JWTServiceの初期化
	jwtService := auth.NewJWTService("your-secret-key")

	// UserのUseCaseとリポジトリの初期化
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Authの''
	authRepo := repository.NewAuthRepository(db)
	authUseCase := usecase.NewAuthUseCase(authRepo, jwtService)

	// Forumの''
	forumRepo := repository.NewForumRepository(db)
	forumUseCase := usecase.NewForumUseCase(forumRepo, jwtService)

	// ルートの設定（依存性注入）
	r := router.SetupRouter(db, jwtService, userUseCase, authUseCase, forumUseCase)
	log.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
