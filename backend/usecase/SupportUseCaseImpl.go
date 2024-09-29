package usecase

import (
	"backend/adapter/repository"
	"backend/domain"
	"backend/infrastructure/auth"
	"errors"
	"time"
)

type SupportUseCaseImpl struct {
	SupportRepository repository.SupportRepository
	UserRepository    repository.UserRepository
	JWTService        *auth.JWTService
}

func NewSupportUseCase(SRepository repository.SupportRepository, URepository repository.UserRepository, jwtService *auth.JWTService) *SupportUseCaseImpl {
	return &SupportUseCaseImpl{
		SupportRepository: SRepository,
		UserRepository:    URepository,
		JWTService:        jwtService,
	}
}

// 新しいサポートリクエストを作成
func (su *SupportUseCaseImpl) NewSupportRequest(supportRequest *domain.SupportRequest) (*domain.SupportRequest, error) {
	// 必要に応じてバリデーションを行う
	if supportRequest.RequestContent == "" {
		return nil, errors.New("request content is required")
	}

	// そのforumに別のsupportRequestがないがチェック
	if supportRequestCheck, err := su.SupportRepository.SelectSupportByForumIDAndStatus(supportRequest.ForumId, domain.StatusInProgress); err != nil {
		return nil, err
	} else if supportRequestCheck != nil {
		return nil, errors.New("support request is exist")
	}

	// 初期状態を設定
	supportRequest.ProgressStatus = domain.StatusInProgress

	// リポジトリを通じてデータベースに保存
	createdSupport, err := su.SupportRepository.CreateSupport(*supportRequest)
	if err != nil {
		return nil, err
	}

	return &createdSupport, nil
}

// サポートリクエストを取得
func (su *SupportUseCaseImpl) GetSupportRequest(forumID int) (*domain.SupportRequest, error) {
	// 指定されたIDのサポートリクエストを取得
	supportRequest, err := su.SupportRepository.SelectSupportByForumIDAndStatus(forumID, domain.StatusInProgress)
	if err != nil {
		return nil, err
	}
	return supportRequest, nil
}

// 取得しにきたUserのDepartmentIDに沿ったサポートリクエストを取得
func (su *SupportUseCaseImpl) GetDepartmentSupportRequests(userID int) ([]*domain.SupportRequest, error) {
	// userIDをactionからもらってDepartmentIDに変換
	user, err := su.UserRepository.SelectByUserID(userID)
	if err != nil {
		return nil, err
	}
	// リポジトリからすべてのサポートリクエストを取得
	supportRequests, err := su.SupportRepository.SelectSupport(user.DepartmentID)
	if err != nil {
		return nil, err
	}

	// []domain.SupportRequest を []*domain.SupportRequest に変換
	var result []*domain.SupportRequest
	for i := range supportRequests {
		result = append(result, &supportRequests[i])
	}

	return result, nil
}

// サポートリクエストを完了状態に更新
func (su *SupportUseCaseImpl) SupportIsComplete(supportID int) (*domain.SupportRequest, error) {
	// 指定されたIDのサポートリクエストを取得
	supportRequest, err := su.SupportRepository.SelectSupportByID(supportID)
	if err != nil {
		return nil, err
	}

	// ステータスを更新
	supportRequest.ProgressStatus = domain.StatusResolved

	// リポジトリを通じてデータベースを更新
	updatedSupport, err := su.SupportRepository.UpdateSupport(*supportRequest)
	if err != nil {
		return nil, err
	}

	return &updatedSupport, nil
}

// サポートリクエストをクローズする
func (su *SupportUseCaseImpl) CloseSupportRequest(forumID int) (*domain.SupportRequest, error) {

	// forumIDとStatusInProgressからsupportRequestを検索する
	supportRequest, err := su.SupportRepository.SelectSupportByForumIDAndStatus(forumID, domain.StatusInProgress)
	if err != nil {
		return nil, err
	}

	// 既に解決またはクローズされている場合はエラーを返す
	if supportRequest.ProgressStatus == domain.StatusResolved || supportRequest.ProgressStatus == domain.StatusClosed {
		return nil, errors.New("このサポートリクエストは既に解決またはクローズされています")
	}

	// ステータスをクローズに更新
	supportRequest.ProgressStatus = domain.StatusClosed
	supportRequest.UpdatedAt = time.Now()

	// リポジトリを通じてデータベースを更新
	updatedSupport, err := su.SupportRepository.CloseSupport(*supportRequest)
	if err != nil {
		return nil, err
	}

	return &updatedSupport, nil
}
