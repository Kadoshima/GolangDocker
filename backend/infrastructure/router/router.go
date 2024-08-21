package router

import (
	"backend/adapter/api/action"
	"backend/usecase"
	"database/sql"
	"net/http"
)

func SetupRouter(db *sql.DB, userUseCase usecase.UserUseCase, authUseCase usecase.AuthUseCase) *http.ServeMux {
	mux := http.NewServeMux()

	// ユーザー作成ハンドラーの設定
	mux.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		action.CreateUserHandler(w, r, userUseCase) // useCaseをハンドラーに渡す
	})

	// ユーザ情報の取得ハンドラーの設定
	mux.HandleFunc("/api/user/InfoGet", func(w http.ResponseWriter, r *http.Request) {
		action.GetUserInfo(w, r, userUseCase) // useCaseをハンドラーに渡す
	})

	// ユーザ情報の取得ハンドラーの設定
	mux.HandleFunc("/api/user/login", func(w http.ResponseWriter, r *http.Request) {
		action.LoginHandler(w, r, authUseCase) // useCaseをハンドラーに渡す
	})

	mux.HandleFunc("/api/user/InfoUpdate", func(w http.ResponseWriter, r *http.Request) {
		action.UpdateUserInfo(w, r, userUseCase) // useCaseをハンドラーに渡す
	})

	return mux
}
