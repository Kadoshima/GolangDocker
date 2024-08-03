package domain

import "time"

type Forums struct {
	Id          int
	Title       string
	Description string
	CreatedById int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
