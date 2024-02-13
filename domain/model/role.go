package model

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        uuid.UUID  `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	Version   int64      `json:"version" db:"version"`
}

type RoleCreate struct {
	ID   uuid.UUID `json:"-"`
	Name string    `json:"name"`
}

type RoleUpdate struct {
	Name string `json:"name"`
}

// Find model

type RoleFilter struct{}

type RoleFind struct {
	Filter RoleFilter
}
