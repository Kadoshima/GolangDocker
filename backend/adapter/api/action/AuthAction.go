package action

import (
	"backend/usecase"
	"encoding/json"
	"net/http"
)

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
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// requestBodyからuserIDを取り出し、string型として扱う
	userIDFloat, ok := requestBody["userID"].(float64)
	if !ok {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}
	// userIDをintに変換
	userID = int(userIDFloat)

	// requestBodyからuserIDを取り出し、string型として扱う
	password, ok2 := requestBody["password"].(string)
	if !ok2 {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}

	// TODO : authUserUseCaseを利用して、ユーザー認証を実装
	auth, err := authUserUseCase.Login(userID, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 成功レスポンス (トークンを返す)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": auth,
	})
}

func Atest(w http.ResponseWriter, r *http.Request, authUserUseCase usecase.AuthUseCase) {
	println("gge")
	return
}
