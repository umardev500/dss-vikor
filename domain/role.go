package domain

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/spk/domain/model"
)

type RoleDelivery interface {
	Create(c *fiber.Ctx) error
}

type RoleUsecase interface {
	Create(ctx context.Context, role model.RoleCreate) model.Response
}

type RoleRepository interface {
	Create(ctx context.Context, role model.RoleCreate) error
}
