package responses

import (
	"github.com/google/uuid"
)

type RiskResp struct {
	ID          uuid.UUID `json:"id"`
	State       string    `json:"state"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   int       `json:"created_at"`
	UpdatedAt   int       `json:"updated_at"`
}
