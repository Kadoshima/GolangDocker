package domain

import "time"

type Forums struct {
	ID          int
	Title       string
	Description string
	CreatedBy   int
	Status      int
	Visibility  int
	Category    string
	NumPosts    int
	Attachments string
	Moderators  []int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
