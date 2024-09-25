package action

import (
	"backend/adapter/api/reqres"
	"backend/usecase"
	"encoding/json"
	"net/http"
)

// LoginHandler godoc
// @Summary      ユーザーのログイン
// @Description  ユーザーがシステムにログインします
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        body  body      models.LoginRequest  true  "ログイン情報"
// @Success      200   {object}  models.LoginSuccessResponse  "成功時のトークン"
// @Failure      400   {object}  models.ErrorResponse         "バリデーションエラー"
// @Router       /api/user/login [post]

func LoginHandler(w http.ResponseWriter, r *http.Request, authUserUseCase usecase.AuthUseCase) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var userIDFloat float64
	var userID int
	var password string

	// リクエストボディを保持するためのmap
	var requestBody map[string]interface{}

	// JSONリクエストボディをデコードしてmapにする
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		reqres.WriteJSONErrorResponse(w, "Invalid request body")
		return
	}

	// requestBodyからuserIDを取り出し、string型として扱う
	userIDFloat, ok := requestBody["userID"].(float64)
	if !ok {
		reqres.WriteJSONErrorResponse(w, "Invalid userID")
		return
	}
	// userIDをintに変換
	userID = int(userIDFloat)

	// requestBodyからuserIDを取り出し、string型として扱う
	password, ok2 := requestBody["password"].(string)
	if !ok2 {
		reqres.WriteJSONErrorResponse(w, "Invalid userID")
		return
	}

	// authUserUseCaseを利用して、ユーザー認証を実装
	auth, err := authUserUseCase.Login(userID, password)
	if err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
		return
	}

	// 成功レスポンス (トークンを返す)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": auth,
	})
}

//func Atest(w http.ResponseWriter, r *http.Request, authUserUseCase usecase.AuthUseCase) {
//
//	// コンテキストからユーザーIDを取得
//	userID, ok := r.Context().Value(middleware.UserContextKey).(int)
//	if !ok {
//		http.Error(w, "User ID not found in context", http.StatusUnauthorized)
//		return
//	}
//
//	// 取得したUserIDを利用して処理を行う
//	println(userID)
//
//	return
//}
