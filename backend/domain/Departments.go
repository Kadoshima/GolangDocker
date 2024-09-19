package domain

import "time"

type Department struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
