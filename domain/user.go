package domain

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/spk/domain/model"
)

type UserDelivery interface {
	Create(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
	Find(*fiber.Ctx) error
	FindOne(*fiber.Ctx) error
	Update(*fiber.Ctx) error
}

type UserUsecase interface {
	Create(context.Context, model.UserCreate) error
	Delete(context.Context, model.UserParams) error
	Find(context.Context, model.UserFind) (model.Response, error)
	FindOne(context.Context, model.UserFind) (model.Response, error)
	Update(context.Context, model.UserUpdate) error
}

type UserRepository interface {
	Create(context.Context, model.UserCreate) error
	Delete(context.Context, model.UserParams) error
	Find(context.Context, model.UserFind) ([]model.User, error)
	FindOne(context.Context, model.UserFind) (model.User, error)
	Update(context.Context, model.UserUpdate) error
}
