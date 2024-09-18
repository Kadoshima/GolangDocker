package action

import (
	"backend/adapter/api/middleware"
	"backend/usecase"
	"net/http"
)

func GetDepartmentAction(w http.ResponseWriter, r *http.Request, useCase usecase.DepartmentUseCase) {
	// コンテキストからユーザーIDを取得
	userID, ok := r.Context().Value(middleware.UserContextKey).(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	// 取得したUserIDを利用して処理を行う
	println(userID)

	return
}
