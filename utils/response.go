package utils

import (
	"github.com/google/uuid"
	"github.com/umardev500/spk/domain/model"
)

func ResponseBuilder(uuid uuid.UUID, code int64, success bool, message interface{}, data interface{}) model.Response {
	response := model.Response{
		ID:      uuid,
		Status:  code,
		Success: success,
		Message: message,
		Data:    data,
	}

	return response
}
