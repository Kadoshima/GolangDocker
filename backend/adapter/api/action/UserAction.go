package action

import (
	"encoding/json"
	"net/http"

	"backend/domain"
	"backend/usecase"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user domain.User

	var UserUsecase usecase.UserUsecase
	if err := UserUsecase.CreateUser(&user); err != nil {
		http.Error(w, "UserUsecase Error", http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"mesage": "User created successfully",
	})

}
