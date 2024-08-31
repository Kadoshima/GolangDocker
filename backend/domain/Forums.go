package domain

import (
	"time"
)

type Forums struct {
	ID          int
	Title       string
	Description string
	CreatedBy   int
	Status      int
	Visibility  int
	Category    string
	NumPosts    int
	Moderators  []int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// 新しいフォーラム作成のファクトリ関数
func NewForum(title string, description string, user *User, visibility int, category string) *Forums {
	return &Forums{
		Title:       title,
		Description: description,
		CreatedBy:   user.ID,
		Status:      1,
		Visibility:  visibility,
		Category:    category,
		NumPosts:    0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
