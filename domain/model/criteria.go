package model

import (
	"time"

	"github.com/google/uuid"
)

type Criteria struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	Version   int64      `json:"version" db:"version"`
}

type CriteriaCreate struct {
	ID   uuid.UUID `json:"-"`
	Name string    `json:"name"`
}

type CriteriaUpdate struct {
	ID   uuid.UUID `json:"-" db:"id" skip:"true"`
	Name string    `json:"name" db:"name"`
}

// Finding
type CriteriaFilter struct{}

type CriteriaFind struct {
	Filter CriteriaFilter
}
