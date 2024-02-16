package model

import (
	"time"

	"github.com/google/uuid"
)

type Alternate struct {
	ID         uuid.UUID  `json:"id" db:"id"`
	Name       string     `json:"name" db:"name"`
	RoleID     string     `json:"role_id" db:"role_id"`
	STR        string     `json:"str" db:"str"`
	Experience int        `json:"experience" db:"experience"`
	DOB        time.Time  `json:"dob" db:"dob"`
	Address    string     `json:"address" db:"address"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	Version    int64      `json:"version" db:"version"`
}

type AlternateCreate struct {
	ID         uuid.UUID `json:"-"`
	Name       string    `json:"name" validate:"required"`
	RoleID     string    `json:"role_id" validate:"required"`
	STR        string    `json:"str" validate:"required"`
	Experience int       `json:"experience" validate:"gte=0"`
	DOB        time.Time `json:"dob" validate:"required"`
	Address    string    `json:"address" validate:"required"`
}

type AlternateUpdate struct {
	ID         uuid.UUID `json:"-" db:"id" skip:"true"`
	Name       string    `json:"name" db:"name" validate:"required"`
	RoleID     string    `json:"role_id" db:"role_id" validate:"required"`
	STR        string    `json:"str" db:"str" validate:"required"`
	Experience int       `json:"experience" db:"experience" validate:"gte=0"`
	DOB        time.Time `json:"dob" db:"dob" validate:"required"`
	Address    string    `json:"address" db:"address" validate:"required"`
}

// Finding
type AlternateFitler struct {
	Name   *string
	RoleID *uuid.UUID
}

type AlternateFind struct {
	Filter   *AlternateFitler
	PageInfo PageInfo
}

var AlternateSorting map[string]string = map[string]string{
	"created_at": "created_at",
	"name":       "name",
	"role":       "role_id",
	"str":        "str",
	"experience": "experience",
	"dob":        "dob",
	"address":    "address",
}
