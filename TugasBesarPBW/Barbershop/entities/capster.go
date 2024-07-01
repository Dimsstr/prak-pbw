package entities

import "time"

type Capster struct {
	ID        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
