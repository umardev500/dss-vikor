package delivery

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/spk/domain"
	"github.com/umardev500/spk/domain/model"
	"github.com/umardev500/spk/utils"
)

type alternateDelivery struct {
	uc domain.AlternateUsecase
	v  *validator.Validate
}

func NewAlternateDelivery(uc domain.AlternateUsecase, r fiber.Router, v *validator.Validate) {
	handler := &alternateDelivery{
		uc: uc,
		v:  v,
	}

	r.Post("/", handler.Create)
}

func (a *alternateDelivery) Create(c *fiber.Ctx) error {
	var alternate model.AlternateCreate
	if err := c.BodyParser(&alternate); err != nil {
		uuid := uuid.New()
		resp := utils.ResponseBuilder(uuid, fiber.StatusBadRequest, false, err.Error(), nil)
		bodyRaw := string(c.Body())
		logData := utils.LogBuilder(uuid, "failed to parse request body", bodyRaw, err)
		log.Error().Msg(logData)
		return c.JSON(resp)
	}

	// Validate the struct
	if err := a.v.Struct(alternate); err != nil {
		uuid := uuid.New()
		resp := utils.ResponseBuilder(uuid, fiber.StatusBadRequest, false, err.Error(), nil)
		bodyRaw := string(c.Body())
		logData := utils.LogBuilder(uuid, fiber.ErrBadRequest.Message, bodyRaw, err)
		log.Error().Msg(logData)
		return c.JSON(resp)
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	resp := a.uc.Create(ctx, alternate)
	return c.JSON(resp)
}
