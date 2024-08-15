package action

import (
	"backend/domain"
	"backend/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request, userUseCase usecase.UserUseCase) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user domain.User

	// JSONリクエストボディからuser構造体をデコードする
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//UserUseCaseを使ってユーザーを作成
	if err := userUseCase.CreateUser(&user); err != nil {
		log.Println(err)
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	// 成功レスポンス
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
}
