package repository

import (
	"backend/domain"
	"database/sql"
	"time"
)

type ForumRepositoryImpl struct {
	db *sql.DB
}

func NewForumRepository(db *sql.DB) *ForumRepositoryImpl {
	return &ForumRepositoryImpl{db: db}
}

func (fr *ForumRepositoryImpl) SelectAllForums() ([]domain.Forums, error) {
	// forumsテーブルから全てのフォーラムを取得
	rows, err := fr.db.Query(
		`SELECT id, title, description, created_by, status, visibility, category, num_posts, created_at, updated_at
		FROM forums`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forums []domain.Forums

	for rows.Next() {
		var forum domain.Forums
		err := rows.Scan(
			&forum.ID, &forum.Title, &forum.Description, &forum.CreatedBy, &forum.Status,
			&forum.Visibility, &forum.Category, &forum.NumPosts, &forum.CreatedAt, &forum.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		// 各フォーラムのモデレーターを取得
		modRows, err := fr.db.Query(
			"SELECT user_id FROM forum_moderators WHERE forum_id = ?", forum.ID,
		)
		if err != nil {
			return nil, err
		}
		defer modRows.Close()

		var moderators []int
		for modRows.Next() {
			var userID int
			if err := modRows.Scan(&userID); err != nil {
				return nil, err
			}
			moderators = append(moderators, userID)
		}
		forum.Moderators = moderators

		forums = append(forums, forum)
	}

	// ループ中にエラーが発生していないか確認
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return forums, nil
}

func (fr *ForumRepositoryImpl) SelectForum(forumID int) (*domain.Forums, error) {
	forum := &domain.Forums{}

	// 基本情報の取得
	err := fr.db.QueryRow(
		`SELECT id, title, description, created_by, status, visibility, category, num_posts, created_at, updated_at
		FROM forums WHERE id = ?`, forumID,
	).Scan(
		&forum.ID, &forum.Title, &forum.Description, &forum.CreatedBy, &forum.Status,
		&forum.Visibility, &forum.Category, &forum.NumPosts, &forum.CreatedAt, &forum.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// モデレーターの取得
	rows, err := fr.db.Query(
		"SELECT user_id FROM forum_moderators WHERE forum_id = ?", forumID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moderators []int
	for rows.Next() {
		var userID int
		if err := rows.Scan(&userID); err != nil {
			return nil, err
		}
		moderators = append(moderators, userID)
	}
	forum.Moderators = moderators

	return forum, nil
}

func (fr *ForumRepositoryImpl) CreateForum(forum *domain.Forums) (*domain.Forums, error) {
	now := time.Now()
	forum.CreatedAt = now
	forum.UpdatedAt = now

	tx, err := fr.db.Begin()
	if err != nil {
		return nil, err
	}

	// forumsテーブルへの挿入
	result, err := tx.Exec(
		`INSERT INTO forums (title, description, created_by, status, visibility, category, num_posts, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		forum.Title, forum.Description, forum.CreatedBy, forum.Status, forum.Visibility,
		forum.Category, forum.NumPosts, forum.CreatedAt, forum.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	forumID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	forum.ID = int(forumID)

	// モデレーターの挿入
	for _, userID := range forum.Moderators {
		_, err := tx.Exec(
			"INSERT INTO forum_moderators (forum_id, user_id) VALUES (?, ?)",
			forum.ID, userID,
		)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return forum, nil
}

func (fr *ForumRepositoryImpl) UpdateForum(forum *domain.Forums) (*domain.Forums, error) {
	forum.UpdatedAt = time.Now()

	tx, err := fr.db.Begin()
	if err != nil {
		return nil, err
	}

	// forumsテーブルの更新
	_, err = tx.Exec(
		`UPDATE forums SET title = ?, description = ?, status = ?, visibility = ?, category = ?, num_posts = ?, updated_at = ?
		WHERE id = ?`,
		forum.Title, forum.Description, forum.Status, forum.Visibility,
		forum.Category, forum.NumPosts, forum.UpdatedAt, forum.ID,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 既存のモデレーターを削除
	_, err = tx.Exec(
		"DELETE FROM forum_moderators WHERE forum_id = ?",
		forum.ID,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 新しいモデレーターを挿入
	for _, userID := range forum.Moderators {
		_, err := tx.Exec(
			"INSERT INTO forum_moderators (forum_id, user_id) VALUES (?, ?)",
			forum.ID, userID,
		)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return forum, nil
}

func (fr *ForumRepositoryImpl) DeleteForum(forum *domain.Forums) error {
	tx, err := fr.db.Begin()
	if err != nil {
		return err
	}

	// モデレーターの削除
	_, err = tx.Exec(
		"DELETE FROM forum_moderators WHERE forum_id = ?",
		forum.ID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// フォーラムの削除
	_, err = tx.Exec(
		"DELETE FROM forums WHERE id = ?",
		forum.ID,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
