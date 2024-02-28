package schemas

import (
	"time"

	"github.com/google/uuid"
)

type Opportunity struct {
	ID        uuid.UUID `gorm:"type:uuid;"`
	Role      string
	Company   string
	Location  string
	Remote    bool
	Link      string
	Salary    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OpportunityResponseBody struct {
	ID        uuid.UUID `json:"id"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    float64   `json:"salary"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
