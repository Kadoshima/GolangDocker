package action

import (
	"backend/adapter/api/middleware"
	"backend/adapter/api/reqres"
	"backend/domain"
	"backend/usecase"
	"encoding/json"
	"net/http"
	"strconv"
)

// CreateSupportRequestAction godoc
// @Summary      サポートリクエストを作成します
// @Description  指定したフォーラムとポストに対するサポートリクエストを追加します
// @Tags         support_request
// @Accept       json
// @Produce      json
// @Param        support_request  body      SupportRequestDto  true  "サポートリクエストデータ"
// @Success      201   {object}  reqres.SupportRequestResponse
// @Failure      400   {object}  map[string]string
// @Router       /api/support_requests [post]
func CreateSupportRequestAction(w http.ResponseWriter, r *http.Request, useCase usecase.SupportUseCase) {

	// メソッドチェック
	if r.Method != http.MethodPost {
		reqres.WriteJSONErrorResponse(w, "Invalid request method")
		return
	}

	// コンテキストからユーザーIDを取得
	userID, ok := r.Context().Value(middleware.UserContextKey).(int)
	if !ok {
		reqres.WriteJSONErrorResponse(w, "User ID not found in context")
		return
	}

	// リクエストボディをデコード
	var supportRequestDto reqres.SupportRequestDto
	if err := json.NewDecoder(r.Body).Decode(&supportRequestDto); err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
		return
	}

	// domain.SupportRequest オブジェクトを作成
	supportRequest := &domain.SupportRequest{
		ForumId:           supportRequestDto.ForumID,
		PostId:            supportRequestDto.PostID,
		RequestContent:    supportRequestDto.RequestContent,
		RequestDepartment: supportRequestDto.RequestDepartment,
		CreatedBy:         userID,
		// ステータスはユースケース内で設定されます
	}

	// ユースケースを呼び出して新しいサポートリクエストを作成
	createdSupportRequest, err := useCase.NewSupportRequest(supportRequest)
	if err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
		return
	}

	// 作成されたサポートリクエストをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(reqres.NewSupportRequestResponse(createdSupportRequest)); err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
		return
	}
}

// CloseSupportRequestAction godoc
// @Summary      サポートリクエストをクローズします
// @Description  指定したフォーラムの進行中のサポートリクエストをクローズします
// @Tags         support_request
// @Accept       json
// @Produce      json
// @Param        forum_id  query     int  true  "フォーラムID"
// @Success      200       {object}  domain.SupportRequest
// @Failure      400       {object}  map[string]string
// @Router       /api/support_requests/close [post]
func CloseSupportRequestAction(w http.ResponseWriter, r *http.Request, useCase usecase.SupportUseCase) {

	// メソッドチェック
	if r.Method != http.MethodPut {
		reqres.WriteJSONErrorResponse(w, "Invalid request method")
		return
	}

	// クエリパラメータから forum_id を取得
	supportIDStr := r.URL.Query().Get("support_id")
	if supportIDStr == "" {
		reqres.WriteJSONErrorResponse(w, "support_id is required")
		return
	}

	// forum_id を整数に変換
	supportID, err := strconv.Atoi(supportIDStr)
	if err != nil {
		reqres.WriteJSONErrorResponse(w, "Invalid forum_id")
		return
	}

	// ユースケースを呼び出してサポートリクエストをクローズ
	closedSupportRequest, err := useCase.CloseSupportRequest(supportID)
	if err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
		return
	}

	// クローズされたサポートリクエストをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(closedSupportRequest); err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
		return
	}
}

// GetDepartmentSupportRequests godoc
// @Summary      サポートリクエスト一覧を取得します
// @Description  すべてのサポートリクエストを取得します
// @Tags         support_request
// @Accept       json
// @Produce      json
// @Success      200   {array}   domain.SupportRequest
// @Failure      404   {object}  map[string]string
// @Router       /api/support_requests [get]
func GetDepartmentSupportRequests(w http.ResponseWriter, r *http.Request, useCase usecase.SupportUseCase) {

	// メソッドチェック
	if r.Method != http.MethodGet {
		reqres.WriteJSONErrorResponse(w, "Invalid request method")
		return
	}

	userID, ok := r.Context().Value(middleware.UserContextKey).(int)
	if !ok {
		reqres.WriteJSONErrorResponse(w, "User ID not found in context")
		return
	}

	// ユースケースを呼び出してサポートリクエストを取得
	supportRequests, err := useCase.GetDepartmentSupportRequests(userID)
	if err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
		return
	}

	// サポートリクエストのリストをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(supportRequests); err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
		return
	}
}

// SupportIsCompleteAction godoc
// @Summary      サポートリクエストを完了状態に更新します
// @Description  指定されたサポートリクエストIDのステータスを解決済みに更新します
// @Tags         support_request
// @Accept       json
// @Produce      json
// @Param        support_id  query     int  true  "サポートリクエストID"
// @Success      200       {object}  domain.SupportRequest
// @Failure      400       {object}  map[string]string
// @Router       /api/support_requests/complete [put]
func SupportIsCompleteAction(w http.ResponseWriter, r *http.Request, useCase usecase.SupportUseCase) {

	// メソッドチェック
	if r.Method != http.MethodPut {
		reqres.WriteJSONErrorResponse(w, "Invalid request method")
		return
	}

	// クエリパラメータから support_id を取得
	supportIDStr := r.URL.Query().Get("support_id")
	if supportIDStr == "" {
		reqres.WriteJSONErrorResponse(w, "support_id is required")
		return
	}

	// support_id を整数に変換
	supportID, err := strconv.Atoi(supportIDStr)
	if err != nil {
		reqres.WriteJSONErrorResponse(w, "Invalid support_id")
		return
	}

	// ユースケースを呼び出してサポートリクエストを完了状態に更新
	updatedSupportRequest, err := useCase.SupportIsComplete(supportID)
	if err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
		return
	}

	// 更新されたサポートリクエストをレスポンスとして返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updatedSupportRequest); err != nil {
		reqres.WriteJSONErrorResponse(w, err.Error())
		return
	}
}
