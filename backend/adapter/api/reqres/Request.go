package reqres

type LoginRequest struct {
	UserID   int    `json:"userID"`
	Password string `json:"password"`
}
