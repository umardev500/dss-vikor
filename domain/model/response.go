package model

import "github.com/google/uuid"

type Response struct {
	ID      uuid.UUID   `json:"id"`
	Status  int64       `json:"status"`
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}
