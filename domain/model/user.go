package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/spk/constants"
)

type UserParams struct {
	ID     uuid.UUID
	UserID uuid.UUID
}

type UserFilter struct{}

type UserFind struct {
	ID      uuid.UUID
	Filters UserFilter
}

type UserUpdate struct {
	Params UserParams
	Data   User
}

type UserCreate struct {
	Params *UserParams
	Data   *UserToCreate
}

// User model to create
type UserToCreate struct {
	ID       uuid.UUID
	Email    string           `json:"email"`
	Password string           `json:"password"`
	Status   constants.Status `json:"status"`
}

// User models
type User struct {
	ID        uuid.UUID        `json:"id" db:"id"`
	Email     string           `json:"email" db:"email"`
	Password  string           `json:"password" db:"password"`
	Status    constants.Status `json:"status" db:"status"`
	CreatedAt time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time       `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time       `json:"deleted_at,omitempty" db:"deleted_at"`
	Version   int64            `json:"version" db:"version"`
}
