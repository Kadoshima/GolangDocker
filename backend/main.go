// @title           ChubuForum API
// @version         1.0
// @description     中部地方のユーザー向けの掲示板アプリ「ChubuForum」のAPI。投稿の作成、閲覧、編集、削除、コメントの投稿などを可能にします。
// @termsOfService  http://your.terms.of.service.url

// @contact.name   あなたの名前
// @contact.url    http://your.contact.url
// @contact.email  your.email@example.com

// @license.name  MIT
// @license.url   http://opensource.org/licenses/MIT

// @host      os3-378-22222.vs.sakura.ne.jp:8000
// @BasePath  /

package main

import (
	"backend/adapter/repository"
	_ "backend/docs"
	"backend/infrastructure"
	"backend/infrastructure/auth"
	"backend/infrastructure/database"
	"backend/infrastructure/router"
	"backend/migrations"
	"backend/service"
	"backend/usecase"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	// 環境変数を呼び出し @1
	endpoint, ok := os.LookupEnv("ENDPOINT")
	if !ok {
		fmt.Println("ENDPOINT is not set")
	}
	accessKeyID, ok := os.LookupEnv("ACCESS_KEY_ID")
	if !ok {
		fmt.Println("ACCESS_KEY_ID is not set")
	}
	secretAccessKey, ok := os.LookupEnv("SECRET_ACCESS_KEY")
	if !ok {
		fmt.Println("SECRET_ACCESS_KEY is not set")
	}
	useSSLStr, ok := os.LookupEnv("USE_SSL")
	if !ok {
		fmt.Println("USE_SSL is not set")
	}
	bucketName, ok := os.LookupEnv("BUCKET_NAME")
	if !ok {
		fmt.Println("BUCKET_NAME is not set")
	}

	//useSSLをboolにconvert
	useSSL, _ := strconv.ParseBool(useSSLStr)

	// DB初期化
	db, err := database.NewDB()
	if err != nil {
		log.Fatalf("Could not connect = %v", err)
	}
	defer db.Close()

	// MinIO初期化
	minioClient, err := infrastructure.NewMinio(endpoint, accessKeyID, secretAccessKey, bucketName, useSSL)
	if err != nil {
		log.Fatalf("Could not connect = %v", err)
	}

	// DBマイグレーション
	migrations.Migrate()

	// JWTServiceの初期化
	jwtService := auth.NewJWTService("your-secret-key")
	minioService := service.NewMinioService(minioClient, bucketName)

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)

	authRepo := repository.NewAuthRepository(db)
	authUseCase := usecase.NewAuthUseCase(authRepo, minioService, jwtService)

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
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
