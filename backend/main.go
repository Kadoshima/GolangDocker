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

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)

	authRepo := repository.NewAuthRepository(db)
	authUseCase := usecase.NewAuthUseCase(authRepo, jwtService)

	forumRepo := repository.NewForumRepository(db)
	forumUseCase := usecase.NewForumUseCase(forumRepo, jwtService)

	postRepo := repository.NewPostRepository(db)
	postUseCase := usecase.NewPostUseCase(postRepo, jwtService)

	courseRepo := repository.NewCourseRepository(db)
	courseUseCase := usecase.NewCourseUseCase(courseRepo, jwtService)

	departmentRepo := repository.NewDepartmentRepository(db)
	departmentUseCase := usecase.NewDepartmentUseCase(departmentRepo, jwtService)

	// ルートの設定（依存性注入）
	r := router.SetupRouter(db, userUseCase, authUseCase, postUseCase, forumUseCase, courseUseCase, departmentUseCase, jwtService)
	log.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
