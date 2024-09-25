package action

import (
	"backend/adapter/api/middleware"
	"backend/adapter/api/response"
	"backend/usecase"
	"encoding/json"
	"net/http"
)

type ForumDTO struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Status      int      `json:"status" validate:"required,oneof=0 1"`     // 例: 0 = 非アクティブ, 1 = アクティブ
	Visibility  int      `json:"visibility" validate:"required,oneof=0 1"` // 例: 0 = 非公開, 1 = 公開
	Category    string   `json:"category" validate:"required"`
	Attachments []string `json:"attachment" validate:"required"`
}

// CreateForumAction CreateForumHandler godoc
// @Summary      フォーラムを作成します
// @Description  新しいフォーラムを追加します
// @Tags         forum
// @Accept       json
// @Produce      json
// @Param        forum  body      domain.Forums  true  "フォーラムデータ"
// @Success      201   {object}  domain.Forums
// @Failure      400   {object}  map[string]string
// @Router       /api/forum/post [post]
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
		forumRequest.Attachments,
	)

	//println(forumRequest.Attachment[1])

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

// GetForumsAction GetForumInfoHandler godoc
// @Summary      フォーラム情報を取得します
// @Description  指定したフォーラムの情報を取得します
// @Tags         forum
// @Accept       json
// @Produce      json
// @Param        forumID  query     int     true  "フォーラムID"
// @Success      200      {object}  domain.Forums
// @Failure      404      {object}  map[string]string
// @Router       /api/forum/get [get]
func GetForumsAction(w http.ResponseWriter, r *http.Request, useCase usecase.ForumUseCase) {
	// メソッドチェック
	if r.Method != http.MethodGet {
		response.WriteJSONErrorResponse(w, "Invalid request method")
		return
	}

	// UseCaseを呼び出してフォーラムのリストを取得
	forums, err := useCase.GetForum()
	if err != nil {
		response.WriteJSONErrorResponse(w, err.Error())
		return
	}

	// フォーラムのリストをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(forums); err != nil {
		response.WriteJSONErrorResponse(w, err.Error())
		return
	}
}
