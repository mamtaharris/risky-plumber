package entities

import (
	"time"

	"github.com/google/uuid"
)

type Risks struct {
	ID          uuid.UUID `json:"id"`
	State       string    `json:"state"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
