package router

import (
	"backend/adapter/api/action"
	"backend/usecase"
	"database/sql"
	"net/http"
)

func SetupRouter(db *sql.DB, userUseCase usecase.UserUseCase) *http.ServeMux {
	mux := http.NewServeMux()

	// ユーザー作成ハンドラーの設定
	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		action.CreateUserHandler(w, r, userUseCase) // useCaseをハンドラーに渡す
	})

	return mux
}
