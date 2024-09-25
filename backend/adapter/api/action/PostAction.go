package action

import (
	"backend/adapter/api/middleware"
	"backend/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

type postDto struct {
	ForumID  int    `json:"forum_id"`
	Content  string `json:"content"`
	Tags     string `json:"tags"`
	ParentID int    `json:"parent_id"`
}

// CreatePostAction CreatePostHandler godoc
// @Summary      ポストを作成します
// @Description  指定したフォーラムにポストを追加します
// @Tags         post
// @Accept       json
// @Produce      json
// @Param        post  body      domain.Post  true  "ポストデータ"
// @Success      201   {object}  domain.Post
// @Failure      400   {object}  map[string]string
// @Router       /api/post/post [post]
func CreatePostAction(w http.ResponseWriter, r *http.Request, useCase usecase.PostUseCase) {

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

	// 引数を受け取ってdomain.Postにdecode
	var postRequest postDto
	if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
		print("hello")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// UseCaseを呼び出して新しい投稿を作成
	createdPost, err := useCase.NewPost(
		postRequest.ForumID,
		userID,
		postRequest.Content,
		postRequest.Tags,
		postRequest.ParentID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 作成された投稿をレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(createdPost); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

// GetPostsAction GetPostsHandler godoc
// @Summary      ポスト一覧を取得します
// @Description  指定したフォーラムのポスト一覧を取得します
// @Tags         post
// @Accept       json
// @Produce      json
// @Param        forum_id  query     int  true  "フォーラムID"
// @Success      200       {array}   domain.Post
// @Failure      404       {object}  map[string]string
// @Router       /api/posts/get [get]
func GetPostsAction(w http.ResponseWriter, r *http.Request, useCase usecase.PostUseCase) {
	// メソッドチェック
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// クエリパラメータから forumID を取得
	forumIDStr := r.URL.Query().Get("forum_id")
	if forumIDStr == "" {
		http.Error(w, "forum_id is required", http.StatusBadRequest)
		return
	}

	// forumID を整数に変換
	forumID, err := strconv.Atoi(forumIDStr)
	if err != nil {
		http.Error(w, "Invalid forum_id", http.StatusBadRequest)
		return
	}

	// UseCaseを呼び出して指定されたフォーラムの投稿を取得
	posts, err := useCase.GetPosts(forumID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 投稿のリストをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
