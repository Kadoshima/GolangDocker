package domain

import "time"

type Report struct {
	Id           int
	ReportedById int
	ForumId      int
	PostId       int
	Reason       string
	ReportedAt   time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
