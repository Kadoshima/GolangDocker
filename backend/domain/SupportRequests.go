package domain

import "time"

type SupportRequestStatus int

const (
	StatusPending    SupportRequestStatus = iota // 0
	StatusInProgress                             // 1
	StatusResolved                               // 2
	StatusClosed                                 // 3
)

type SupportRequest struct {
	Id                int
	ForumId           int
	PostId            int
	RequestContent    string
	RequestDepartment int
	CreatedBy         int
	ProgressStatus    SupportRequestStatus
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
