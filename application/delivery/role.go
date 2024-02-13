package delivery

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/spk/domain"
	"github.com/umardev500/spk/domain/model"
	"github.com/umardev500/spk/utils"
)

type roleDelivery struct {
	uc domain.RoleUsecase
}

func NewRoleDelivery(uc domain.RoleUsecase, r fiber.Router) {
	handler := &roleDelivery{
		uc: uc,
	}

	r.Post("/", handler.Create)
}

func (r *roleDelivery) Create(c *fiber.Ctx) error {
	var role model.RoleCreate

	if err := c.BodyParser(&role); err != nil {
		uuid := uuid.New()
		resp := utils.ResponseBuilder(uuid, fiber.StatusBadRequest, false, err.Error(), nil)
		bodyRaw := string(c.Body())
		logData := utils.LogBuilder(uuid, "failed to parse request body", bodyRaw, err)
		log.Error().Msg(logData)

		return c.JSON(resp)
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	resp := r.uc.Create(ctx, role)
	return c.JSON(resp)
}
