package action

import (
	"backend/domain"
	"backend/usecase"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

func GetUserInfo(w http.ResponseWriter, r *http.Request, userUseCase usecase.UserUseCase) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var userIDStr string
	var userID int

	userIDStr = r.URL.Query().Get("userID")
	if userIDStr == "" {
		http.Error(w, "userID is required", http.StatusBadRequest)
		return
	}

	// userIDをintに変換
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid userID", http.StatusBadRequest)
		return
	}

	//UserUseCaseを使ってユーザーを作成
	user, err := userUseCase.UserInfoGet(userID)
	if err != nil {
		http.Error(w, "No user", http.StatusInternalServerError)
		return
	}

	// 成功レスポンス
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":        "I20400",
		"message":       "User Info Get successfully",
		"id":            user.ID,
		"student_id":    user.StudentID,
		"nickname":      user.Nickname,
		"email":         user.Email,
		"department_id": user.DepartmentID,
		"course_id":     user.CourseID,
	})
}

func UpdateUserInfo(w http.ResponseWriter, r *http.Request, userUseCase usecase.UserUseCase) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := userUseCase.UserInfoUpdate(&user); err != nil {
		log.Println(err)
		http.Error(w, "Could not update user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) ///////////////////////////////////////////
	json.NewEncoder(w).Encode(map[string]string{
		// DBと比較していないので，DBと同じでも表示される
		"message": "User information updated successfully",
	})
}
