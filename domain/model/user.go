package model

import (
	"database/sql"
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
	ID        uuid.UUID        `json:"id"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	Status    constants.Status `json:"status"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt sql.NullTime     `json:"updated_at"`
	DeletedAt sql.NullTime     `json:"deleted_at"`
}
