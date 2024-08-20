package action

import (
	"backend/usecase"
	"encoding/json"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, userUseCase usecase.UserUseCase, authUserUseCase usecase.AuthUseCase) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// リクエストボディを保持するためのmap
	var requestBody map[string]interface{}

	// JSONリクエストボディをデコードしてmapにする
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TODO : userUseCaseを利用してUser情報を取得

	// TODO : userUseCaseやauthUserUseCaseを利用して、ユーザー認証を実装

	// // ここでは簡易的な認証の例を示します。
	// if username != "testuser" || password != "password" {
	// 	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	// 	return
	// }

	// // 認証が成功したらJWTを生成
	// token, err := auth.GenerateJWT(username)
	// if err != nil {
	// 	http.Error(w, "Could not generate token", http.StatusInternalServerError)
	// 	return
	// }

	// 成功レスポンス (トークンを返す)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
