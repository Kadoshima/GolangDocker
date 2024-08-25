package router

import (
	"backend/adapter/api/action"
	"backend/adapter/api/middleware"
	"backend/infrastructure/auth"
	"backend/usecase"
	"database/sql"
	"net/http"
)

func SetupRouter(db *sql.DB, userUseCase usecase.UserUseCase, authUseCase usecase.AuthUseCase, jwtService *auth.JWTService) *http.ServeMux {
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

	// 認証が必要なルートにミドルウェアを適用
	mux.Handle("/api/atest", middleware.JWTMiddleware(jwtService)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		action.Atest(w, r, authUseCase) // useCaseをハンドラーに渡す
	})))

	return mux
}
