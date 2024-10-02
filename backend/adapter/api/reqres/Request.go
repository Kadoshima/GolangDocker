package reqres

type LoginRequest struct {
	UserID   int    `json:"userID"`
	Password string `json:"password"`
}

type SupportRequestDto struct {
	ForumID           int    `json:"forum_id" validate:"required"`
	PostID            int    `json:"post_id" validate:"required"`
	RequestContent    string `json:"request_content" validate:"required"`
	RequestDepartment int    `json:"request_department" validate:"required"`
}
