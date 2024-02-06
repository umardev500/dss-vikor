package utils

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func StructToJson(data interface{}) string {
	res, err := json.Marshal(data)
	if err != nil {
		logData := LogBuilder(uuid.New(), "failed to marshal json", "", err)
		log.Error().Msg(logData)
	}

	return string(res)
}
