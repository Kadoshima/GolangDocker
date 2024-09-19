package repository

import (
	"backend/domain"
	"database/sql"
	"time"
)

type PostRepositoryImpl struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepositoryImpl {
	return &PostRepositoryImpl{db: db}
}

func (pr *PostRepositoryImpl) SelectPost(forumID int) (*domain.Post, error) {
	post := &domain.Post{}
	var parentID sql.NullInt64

	err := pr.db.QueryRow(
		`SELECT id, forum_id, user_id, content, tags, status, parent_id, created_at, updated_at
		FROM posts WHERE id = ?`, forumID,
	).Scan(
		&post.ID, &post.ForumId, &post.UserId, &post.Content, &post.Tags,
		&post.Status, &parentID, &post.CreatedAt, &post.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	if parentID.Valid {
		val := int(parentID.Int64)
		post.ParentId = val
	}

	// 添付ファイルの取得
	//attachments, err := pr.getAttachments(post.ID)
	//if err != nil {
	//	return nil, err
	//}
	//post.Attachments = attachments

	return post, nil
}

func (pr *PostRepositoryImpl) CreatePost(post *domain.Post) (*domain.Post, error) {
	now := time.Now()
	post.CreatedAt = now
	post.UpdatedAt = now

	tx, err := pr.db.Begin()
	if err != nil {
		return nil, err
	}

	var parentID interface{}

	// postsテーブルへの挿入
	result, err := tx.Exec(
		`INSERT INTO posts (forum_id, user_id, content, tags, status, parent_id, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		post.ForumId, post.UserId, post.Content, post.Tags, post.Status, parentID, post.CreatedAt, post.UpdatedAt,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	post.ID = int(postID)

	// 添付ファイルの挿入
	for _, attachment := range post.Attachments {
		_, err := tx.Exec(
			"INSERT INTO post_attachments (post_id, attachment) VALUES (?, ?)",
			post.ID, attachment,
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

	return post, nil
}

func (pr *PostRepositoryImpl) UpdatePost(post *domain.Post) (*domain.Post, error) {
	post.UpdatedAt = time.Now()

	tx, err := pr.db.Begin()
	if err != nil {
		return nil, err
	}

	var parentID interface{}

	// postsテーブルの更新
	_, err = tx.Exec(
		`UPDATE posts SET forum_id = ?, user_id = ?, content = ?, tags = ?, status = ?, parent_id = ?, updated_at = ?
		WHERE id = ?`,
		post.ForumId, post.UserId, post.Content, post.Tags, post.Status, parentID, post.UpdatedAt, post.ID,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 既存の添付ファイルを削除
	_, err = tx.Exec(
		"DELETE FROM post_attachments WHERE post_id = ?", post.ID,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// 新しい添付ファイルを挿入
	for _, attachment := range post.Attachments {
		_, err := tx.Exec(
			"INSERT INTO post_attachments (post_id, attachment) VALUES (?, ?)",
			post.ID, attachment,
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

	return post, nil
}

func (pr *PostRepositoryImpl) DeletePost(postID int) error {
	tx, err := pr.db.Begin()
	if err != nil {
		return err
	}

	// 添付ファイルの削除
	//_, err = tx.Exec(
	//	"DELETE FROM post_attachments WHERE post_id = ?", postID,
	//)
	//if err != nil {
	//	tx.Rollback()
	//	return err
	//}

	// ポストの削除
	_, err = tx.Exec(
		"DELETE FROM posts WHERE id = ?", postID,
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
