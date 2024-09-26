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
// @Param        body  body      reqres.LoginRequest  true  "ログイン情報"
// @Success      200   {object}  reqres.LoginSuccessResponse  "成功時のトークン"
// @Failure      400   {object}  reqres.ErrorResponse         "バリデーションエラー"
// @Router       /api/user/login [post]
func LoginHandler(w http.ResponseWriter, r *http.Request, authUserUseCase usecase.AuthUseCase) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var password string

	// リクエストボディを保持するためのmap
	var requestBody map[string]string

	// JSONリクエストボディをデコードしてmapにする
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		reqres.WriteJSONErrorResponse(w, "Invalid request body")
		return
	}

	// requestBodyからuserIDを取り出し、string型として扱う
	studentID, ok := requestBody["student_id"]

	if !ok {
		reqres.WriteJSONErrorResponse(w, "Invalid student_id")
		return
	}

	// requestBodyからpassを取り出し、string型として扱う
	password, ok2 := requestBody["password"]
	if !ok2 {
		reqres.WriteJSONErrorResponse(w, "Invalid password")
		return
	}

	//studentIDからuserIDの取得
	userID, err := authUserUseCase.GetUserIDByStudentID(studentID)
	if err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
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
