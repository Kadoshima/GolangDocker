package reqres

type LoginRequest struct {
	UserID   int    `json:"userID"`
	Password string `json:"password"`
}

type SupportRequestDto struct {
	ForumID        int    `json:"forum_id"`
	PostID         int    `json:"post_id"`
	RequestContent string `json:"request_content"`
}
