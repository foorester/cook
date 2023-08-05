package transport

import "time"

type (
	GetBookRes struct {
		ServiceRes
		ID          string
		Name        string
		Description string
		OwnerID     string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
