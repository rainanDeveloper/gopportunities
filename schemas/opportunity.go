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
