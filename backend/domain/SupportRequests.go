package domain

import "time"

type SupportRequest struct {
	Id             int
	ForumId        int
	PostId         int
	RequestContent string
	IsResolved     bool
	IdInProgress   bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
