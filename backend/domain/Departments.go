package domain

import "time"

type Department struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
