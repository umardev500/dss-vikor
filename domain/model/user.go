package model

import (
	"database/sql"
	"time"
)

type UserStatus int64

const (
	UserStatusActive UserStatus = iota
	UserStatusInactive
)

type UserParams struct {
	ID     string
	UserID string
}

type UserFilter struct{}

type UserFind struct {
	ID      string
	Filters UserFilter
}

type UserUpdate struct {
	Params UserParams
	Data   User
}

// User models
type User struct {
	ID        string       `json:"id"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	Status    UserStatus   `json:"status"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
