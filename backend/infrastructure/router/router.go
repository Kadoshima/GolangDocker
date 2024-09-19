package router

import (
	"backend/adapter/api/action"
	"backend/adapter/api/middleware"
	"backend/infrastructure/auth"
	"backend/usecase"
	"database/sql"
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
	jwtService *auth.JWTService,
) *http.ServeMux {

	mux := http.NewServeMux()

	// ユーザー作成ハンドラーの設定
	mux.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		action.CreateUserHandler(w, r, userUseCase) // useCaseをハンドラーに渡す
	})

	mux.HandleFunc("/api/user/InfoGet", func(w http.ResponseWriter, r *http.Request) {
		action.GetUserInfo(w, r, userUseCase) // useCaseをハンドラーに渡す
	})

	mux.HandleFunc("/api/user/login", func(w http.ResponseWriter, r *http.Request) {
		action.LoginHandler(w, r, authUseCase) // useCaseをハンドラーに渡す
	})

	mux.HandleFunc("/api/user/InfoUpdate", func(w http.ResponseWriter, r *http.Request) {
		action.UpdateUserInfo(w, r, userUseCase) // useCaseをハンドラーに渡す
	})

	// 新しい掲示板(Forum)を作成するAPI
	mux.Handle("/api/forum", middleware.JWTMiddleware(jwtService)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action.CreateForumAction(w, r, forumUseCase) // useCaseをハンドラーに渡す
	})))

	// 新しいpostを作成するAPI
	mux.Handle("/api/post", middleware.JWTMiddleware(jwtService)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action.PostAction(w, r, postUseCase) // useCaseをハンドラーに渡す
	})))

	// 全てのCourse情報の取得
	mux.HandleFunc("/api/allCourse", func(w http.ResponseWriter, r *http.Request) {
		action.GetAllCourseInfoAction(w, r, courseUseCase) // useCaseをハンドラーに渡す
	})

	// Course情報の取得
	mux.HandleFunc("/api/Course", func(w http.ResponseWriter, r *http.Request) {
		action.GetCourseInfoAction(w, r, courseUseCase) // useCaseをハンドラーに渡す
	})

	// Department情報の取得
	mux.HandleFunc("/api/allDepartments", func(w http.ResponseWriter, r *http.Request) {
		action.GetAllDepartmentAction(w, r, departmentUseCase) // useCaseをハンドラーに渡す
	})

	// Department情報の取得
	mux.HandleFunc("/api/departments", func(w http.ResponseWriter, r *http.Request) {
		action.GetDepartmentAction(w, r, departmentUseCase) // useCaseをハンドラーに渡す
	})

	return mux
}
