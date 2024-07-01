package entities

import "time"

type Service struct {
	ID        int
	Name      string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
