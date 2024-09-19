package action

import (
	"backend/adapter/api/middleware"
	"backend/usecase"
	"encoding/json"
	"net/http"
)

type ForumDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      int    `json:"status" validate:"required,oneof=0 1"`     // 例: 0 = 非アクティブ, 1 = アクティブ
	Visibility  int    `json:"visibility" validate:"required,oneof=0 1"` // 例: 0 = 非公開, 1 = 公開
	Category    string `json:"category" validate:"required"`
}

func CreateForumAction(w http.ResponseWriter, r *http.Request, useCase usecase.ForumUseCase) {
	// メソッドチェック
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// コンテキストからユーザーIDを取得
	userID, ok := r.Context().Value(middleware.UserContextKey).(int)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
		return
	}

	println(userID)

	// リクエストボディをForumDTOにデコード
	var forumRequest ForumDTO
	if err := json.NewDecoder(r.Body).Decode(&forumRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// UseCaseを呼び出して新しいフォーラムを作成
	createdForum, err := useCase.CreateForum(
		forumRequest.Title,
		forumRequest.Description,
		userID,
		forumRequest.Status,
		forumRequest.Visibility,
		forumRequest.Category,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 作成されたフォーラムをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(createdForum); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}
