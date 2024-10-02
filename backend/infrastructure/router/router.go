package router

import (
	"backend/adapter/api/action"
	"backend/adapter/api/middleware"
	"backend/infrastructure/auth"
	"backend/usecase"
	"database/sql"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func SetupRouter(db *sql.DB,
	// 各種UseCaseを定義
	userUseCase usecase.UserUseCase,
	authUseCase usecase.AuthUseCase,
	postUseCase usecase.PostUseCase,
	forumUseCase usecase.ForumUseCase,
	courseUseCase usecase.CourseUseCase,
	departmentUseCase usecase.DepartmentUseCase,
	supportUseCase usecase.SupportUseCase,
	jwtService *auth.JWTService,
) *http.ServeMux {

	mux := http.NewServeMux()

	// swaggerの設定
	mux.Handle("/doc/", httpSwagger.WrapHandler)

	// ユーザー作成ハンドラーの設定
	mux.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		action.CreateUserHandler(w, r, userUseCase) // useCaseをハンドラーに渡す
	})

	mux.HandleFunc("/api/user/get", func(w http.ResponseWriter, r *http.Request) {
		action.GetUserInfo(w, r, userUseCase) // useCaseをハンドラーに渡す
	})

	mux.HandleFunc("/api/user/login", func(w http.ResponseWriter, r *http.Request) {
		action.LoginHandler(w, r, authUseCase) // useCaseをハンドラーに渡す
	})

	mux.HandleFunc("/api/user/update", func(w http.ResponseWriter, r *http.Request) {
		action.UpdateUserInfo(w, r, userUseCase) // useCaseをハンドラーに渡す
	})

	// 新しい掲示板(Forum)を作成するAPI
	mux.Handle("/api/forum/post", middleware.JWTMiddleware(jwtService)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action.CreateForumAction(w, r, forumUseCase) // useCaseをハンドラーに渡す
	})))

	// 掲示板情報を取得する
	mux.Handle("/api/forum/get", middleware.JWTMiddleware(jwtService)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action.GetForumsAction(w, r, forumUseCase) // useCaseをハンドラーに渡す
	})))

	// 新しいpostを作成するAPI
	mux.Handle("/api/post/post", middleware.JWTMiddleware(jwtService)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action.CreatePostAction(w, r, postUseCase) // useCaseをハンドラーに渡す
	})))

	// forumに対応するpost群を取る
	mux.Handle("/api/posts/get", middleware.JWTMiddleware(jwtService)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action.GetPostsAction(w, r, postUseCase) // useCaseをハンドラーに渡す
	})))

	// 全てのCourse情報の取得
	mux.HandleFunc("/api/courses/get", func(w http.ResponseWriter, r *http.Request) {
		action.GetAllCourseInfoAction(w, r, courseUseCase) // useCaseをハンドラーに渡す
	})

	// Course情報の取得
	mux.HandleFunc("/api/course", func(w http.ResponseWriter, r *http.Request) {
		action.GetCourseInfoAction(w, r, courseUseCase) // useCaseをハンドラーに渡す
	})

	// Department情報の取得
	mux.HandleFunc("/api/departments/get", func(w http.ResponseWriter, r *http.Request) {
		action.GetAllDepartmentAction(w, r, departmentUseCase) // useCaseをハンドラーに渡す
	})

	// Department情報の取得
	mux.HandleFunc("/api/department/get", func(w http.ResponseWriter, r *http.Request) {
		action.GetDepartmentAction(w, r, departmentUseCase) // useCaseをハンドラーに渡す
	})

	// support系
	mux.Handle("/api/support/post", middleware.JWTMiddleware(jwtService)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action.CreateSupportRequestAction(w, r, supportUseCase) // useCaseをハンドラーに渡す
	})))

	// 自分の所属する学部に当てられたsupportを全てもら
	mux.Handle("/api/support/get", middleware.JWTMiddleware(jwtService)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action.GetDepartmentSupportRequests(w, r, supportUseCase) // useCaseをハンドラーに渡す
	})))

	mux.Handle("/api/support/complete", middleware.JWTMiddleware(jwtService)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action.SupportIsCompleteAction(w, r, supportUseCase) // useCaseをハンドラーに渡す
	})))

	mux.Handle("/api/support/close", middleware.JWTMiddleware(jwtService)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action.CloseSupportRequestAction(w, r, supportUseCase) // useCaseをハンドラーに渡す
	})))

	return mux
}
