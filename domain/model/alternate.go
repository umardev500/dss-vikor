package model

import (
	"time"

	"github.com/google/uuid"
)

type Alternate struct {
	ID         uuid.UUID  `json:"id"`
	Name       string     `json:"name"`
	RoleID     string     `json:"role_id"`
	STR        string     `json:"str"`
	Experience int        `json:"experience"`
	DOB        time.Time  `json:"dob"`
	Address    string     `json:"address"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	Version    int        `json:"version"`
}

type AlternateCreate struct {
	ID         uuid.UUID `json:"-"`
	Name       string    `json:"name" validate:"required"`
	RoleID     string    `json:"role_id" validate:"required"`
	STR        string    `json:"str"`
	Experience int       `json:"experience" validate:"gte=0"`
	DOB        time.Time `json:"dob" validate:"required"`
	Address    string    `json:"address" validate:"required"`
}
