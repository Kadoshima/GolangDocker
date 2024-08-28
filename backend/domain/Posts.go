package domain

import "time"

type Post struct {
	ID          int
	ForumId     int
	UserId      int
	Content     string
	Tags        string
	Status      int
	ParentId    *int
	Attachments []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
