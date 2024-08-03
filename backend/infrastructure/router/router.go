package router

import (
	"backend/adapter/api/action"
	"net/http"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/users", action.CreateUserHandler) // POST /api/users でユーザー作成

	return mux
}
