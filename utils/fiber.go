package utils

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/spk/domain/model"
)

func GetRawBodySingleLine(c *fiber.Ctx) (res string, hndl, err error) {
	bodyRaw := string(c.Body())
	var data map[string]interface{}
	uid := uuid.New()
	resp := model.Response{
		ID:      uid,
		Status:  fiber.StatusInternalServerError,
		Success: false,
		Message: fiber.ErrInternalServerError.Message,
	}

	if err = json.Unmarshal([]byte(bodyRaw), &data); err != nil {
		logData := LogBuilder(uid, "failed to unmarshal json body", bodyRaw, err)
		log.Error().Msg(logData)
		return res, c.JSON(resp), err
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		logData := LogBuilder(uid, "failed to marshal json body", bodyRaw, err)
		log.Error().Msg(logData)
		return res, c.JSON(resp), err
	}

	res = string(jsonData)

	return
}

func ParseUUID(c *fiber.Ctx) (*uuid.UUID, error) {
	var id = c.Params("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		userMsg := "please provide valid uuid"
		debugID := uuid.New()
		resp := ResponseBuilder(debugID, fiber.StatusBadRequest, false, userMsg, nil)
		bodyRaw := string(c.BodyRaw())
		logData := LogBuilder(debugID, "failed to parse uuid", bodyRaw, err)
		log.Error().Msg(logData)

		return nil, c.JSON(resp)
	}

	return &uid, nil
}
