package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/spk/utils"
)

func ParseUUID(id string, c *fiber.Ctx) (*uuid.UUID, error) {
	uid, err := uuid.Parse(id)

	if err != nil {
		userMsg := "please provide valid uuid"
		debugID := uuid.New()
		resp := utils.ResponseBuilder(debugID, fiber.StatusBadRequest, false, userMsg, nil)
		bodyRaw := string(c.BodyRaw())
		logData := utils.LogBuilder(debugID, "failed to parse uuid", bodyRaw, err)
		log.Error().Msg(logData)

		return nil, c.JSON(resp)
	}

	return &uid, nil
}
