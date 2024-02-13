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
}
