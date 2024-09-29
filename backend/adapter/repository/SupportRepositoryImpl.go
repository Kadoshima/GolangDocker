package repository

import (
	"backend/domain"
	"database/sql"
	"time"
)

type SupportRepositoryImpl struct {
	db *sql.DB
}

func NewSupportRepository(db *sql.DB) *SupportRepositoryImpl {
	return &SupportRepositoryImpl{db: db}
}

// サポートリクエストの新規作成
func (sr *SupportRepositoryImpl) CreateSupport(support domain.SupportRequest) (domain.SupportRequest, error) {
	now := time.Now()
	support.CreatedAt = now
	support.UpdatedAt = now

	result, err := sr.db.Exec(
		`INSERT INTO support_requests (forum_id, post_id, request_content, progress_status, created_at, updated_at)
		 VALUES (?, ?, ?, ?, ?, ?)`,
		support.ForumId, support.PostId, support.RequestContent, support.ProgressStatus, support.CreatedAt, support.UpdatedAt,
	)
	if err != nil {
		return support, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return support, err
	}
	support.Id = int(id)
	return support, nil
}

// サポートリクエストの取得
func (sr *SupportRepositoryImpl) SelectSupport() ([]domain.SupportRequest, error) {
	rows, err := sr.db.Query(
		`SELECT id, forum_id, post_id, request_content, progress_status, created_at, updated_at FROM support_requests`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var supports []domain.SupportRequest
	for rows.Next() {
		var support domain.SupportRequest
		if err := rows.Scan(
			&support.Id,
			&support.ForumId,
			&support.PostId,
			&support.RequestContent,
			&support.ProgressStatus,
			&support.CreatedAt,
			&support.UpdatedAt,
		); err != nil {
			return nil, err
		}
		supports = append(supports, support)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return supports, nil
}

// サポートリクエストの更新
func (sr *SupportRepositoryImpl) UpdateSupport(support domain.SupportRequest) (domain.SupportRequest, error) {
	support.UpdatedAt = time.Now()

	_, err := sr.db.Exec(
		`UPDATE support_requests SET forum_id = ?, post_id = ?, request_content = ?, progress_status = ?, updated_at = ? WHERE id = ?`,
		support.ForumId, support.PostId, support.RequestContent, support.ProgressStatus, support.UpdatedAt, support.Id,
	)
	if err != nil {
		return support, err
	}
	return support, nil
}

// サポートリクエストの削除
func (sr *SupportRepositoryImpl) DeleteSupport(support domain.SupportRequest) error {
	_, err := sr.db.Exec(
		`DELETE FROM support_requests WHERE id = ?`,
		support.Id,
	)
	return err
}

// サポートリクエストの完了処理
func (sr *SupportRepositoryImpl) CompleteSupport(support domain.SupportRequest) (domain.SupportRequest, error) {
	support.UpdatedAt = time.Now()

	_, err := sr.db.Exec(
		`UPDATE support_requests SET progress_status = ?, updated_at = ? WHERE id = ?`,
		support.ProgressStatus, support.UpdatedAt, support.Id,
	)
	if err != nil {
		return support, err
	}
	return support, nil
}

// 指定されたIDのサポートリクエストを取得
func (sr *SupportRepositoryImpl) SelectSupportByID(supportID int) (*domain.SupportRequest, error) {
	support := &domain.SupportRequest{}
	err := sr.db.QueryRow(
		`SELECT id, forum_id, post_id, request_content, progress_status, created_at, updated_at
		 FROM support_requests WHERE id = ?`,
		supportID,
	).Scan(
		&support.Id,
		&support.ForumId,
		&support.PostId,
		&support.RequestContent,
		&support.ProgressStatus,
		&support.CreatedAt,
		&support.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return support, nil
}

// サポートリクエストをクローズする
func (sr *SupportRepositoryImpl) CloseSupport(support domain.SupportRequest) (domain.SupportRequest, error) {
	support.ProgressStatus = domain.StatusClosed
	support.UpdatedAt = time.Now()

	_, err := sr.db.Exec(
		`UPDATE support_requests SET progress_status = ?, updated_at = ? WHERE id = ?`,
		support.ProgressStatus, support.UpdatedAt, support.Id,
	)
	if err != nil {
		return support, err
	}
	return support, nil
}

// 指定されたforumIDとstatusを持つサポートリクエストを取得
func (sr *SupportRepositoryImpl) SelectSupportByForumIDAndStatus(forumID int, status domain.SupportRequestStatus) (*domain.SupportRequest, error) {
	support := &domain.SupportRequest{}
	err := sr.db.QueryRow(
		`SELECT id, forum_id, post_id, request_content, progress_status, created_at, updated_at
		 FROM support_requests WHERE forum_id = ? AND progress_status = ?`,
		forumID, status,
	).Scan(
		&support.Id,
		&support.ForumId,
		&support.PostId,
		&support.RequestContent,
		&support.ProgressStatus,
		&support.CreatedAt,
		&support.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return support, nil
}
