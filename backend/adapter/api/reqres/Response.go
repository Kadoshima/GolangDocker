package reqres

import (
	"backend/domain"
	"encoding/json"
	"net/http"
	"time"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type LoginSuccessResponse struct {
	Token string `json:"token"`
}

type SupportRequestResponse struct {
	ID             int    `json:"id"`
	ForumID        int    `json:"forum_id"`
	PostID         int    `json:"post_id"`
	RequestContent string `json:"request_content"`
	CreatedBy      int    `json:"created_by"`
	ProgressStatus int    `json:"progress_status"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

func WriteJSONErrorResponse(w http.ResponseWriter, errorMessage string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(ErrorResponse{errorMessage})
}

func NewSupportRequestResponse(supportRequest *domain.SupportRequest) *SupportRequestResponse {
	return &SupportRequestResponse{
		ID:             supportRequest.Id,
		ForumID:        supportRequest.ForumId,
		PostID:         supportRequest.PostId,
		RequestContent: supportRequest.RequestContent,
		CreatedBy:      supportRequest.CreatedBy,
		ProgressStatus: int(supportRequest.ProgressStatus),
		CreatedAt:      supportRequest.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      supportRequest.UpdatedAt.Format(time.RFC3339),
	}
}
