package domain

import "time"

type Post struct {
	ID        int
	ForumId   int
	UserId    int
	Content   string
	Tags      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
