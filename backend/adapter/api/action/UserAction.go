package action

import (
	"backend/adapter/api/middleware"
	"backend/adapter/api/reqres"
	"backend/domain"
	"backend/usecase"
	"encoding/json"
	"log"
	"net/http"
)

// CreateUserHandler godoc
// @Summary      ユーザーを作成します
// @Description  新しいユーザーをシステムに追加します
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body      domain.User  true  "ユーザーデータ"
// @Success      201   {object}  domain.User
// @Failure      400   {object}  map[string]string
// @Router       /api/user [post]
func CreateUserHandler(w http.ResponseWriter, r *http.Request, userUseCase usecase.UserUseCase) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user domain.User

	// JSONリクエストボディからuser構造体をデコードする
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		reqres.WriteJSONErrorResponse(w, "Invalid request body")
		return
	}

	//UserUseCaseを使ってユーザーを作成
	if err := userUseCase.CreateUser(&user); err != nil {
		log.Println(err)
		reqres.WriteJSONErrorResponse(w, "Could not create user")
		return
	}

	// 成功レスポンス
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User created successfully",
	})
}

// GetUserInfo GetUserInfoHandler godoc
// @Summary      ユーザー情報を取得します
// @Description  指定したユーザーの情報を取得します
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        userID  query     int     true  "ユーザーID"
// @Success      200     {object}  domain.User
// @Failure      404     {object}  map[string]string
// @Router       /api/user/get [get]
func GetUserInfo(w http.ResponseWriter, r *http.Request, userUseCase usecase.UserUseCase) {
	if r.Method != http.MethodGet {
		reqres.WriteJSONErrorResponse(w, "Invalid request method")
		return
	}

	// コンテキストからユーザーIDを取得
	userID, ok := r.Context().Value(middleware.UserContextKey).(int)
	if !ok {
		reqres.WriteJSONErrorResponse(w, "User ID not found in context")
		return
	}

	//UserInfoGetを使ってユーザー情報取得
	user, err := userUseCase.UserInfoGet(userID)
	if err != nil {
		reqres.WriteJSONErrorResponse(w, "No user")
		return
	}

	// 成功レスポンス
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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

// UpdateUserInfo GetUserInfoHandler godoc
// @Summary      ユーザー情報を修正します
// @Description  指定したユーザーの情報を修正します
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        userID  query     int     true  "ユーザーID"
// @Success      200     {object}  domain.User
// @Failure      404     {object}  map[string]string
// @Router       /api/user/update [post]
func UpdateUserInfo(w http.ResponseWriter, r *http.Request, userUseCase usecase.UserUseCase) {
	if r.Method != http.MethodPut {
		reqres.WriteJSONErrorResponse(w, "Invalid request method")
		return
	}

	var user domain.User
	// JSONリクエストボディからuser構造体をデコードする
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		reqres.WriteJSONErrorResponse(w, "Invalid request body")
		return
	}
	// UserInfoUpdateを使ってユーザー情報更新
	if err := userUseCase.UserInfoUpdate(&user); err != nil {
		reqres.WriteJSONErrorResponse(w, "Could not update user")
		return
	}
	// 成功レスポンス
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		// DBと比較していないので，DBと同じでも表示される
		"message": "User information updated successfully",
	})
}
