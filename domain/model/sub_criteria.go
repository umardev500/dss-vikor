package model

import (
	"time"

	"github.com/google/uuid"
)

type SubCriteria struct {
	ID         uuid.UUID  `json:"id" db:"id"`
	CriteriaID uuid.UUID  `json:"criteria_id" db:"criteria_id"`
	Name       string     `json:"name" db:"name"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	Version    int64      `json:"version" db:"version"`
}

type SubCriteriaCreate struct {
	ID         uuid.UUID `json:"-"`
	CriteriaID uuid.UUID `json:"criteria_id"`
	Name       string    `json:"name"`
}

type SubCriteriaUpdate struct {
	ID         uuid.UUID `json:"-" db:"id" skip:"true"`
	CriteriaID uuid.UUID `json:"criteria_id" db:"criteria_id"`
	Name       string    `json:"name" db:"name"`
}

// Finding
type SubCriteriaFilter struct{}

type SubCriteriaFind struct {
	Filter SubCriteriaFilter
}
