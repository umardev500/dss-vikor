package delivery

import (
	"context"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/spk/domain"
	"github.com/umardev500/spk/domain/model"
)

type userDelivery struct {
	uc domain.UserUsecase
}

func NewUserDelivery(uc domain.UserUsecase) domain.UserDelivery {
	return &userDelivery{
		uc: uc,
	}
}

// Create implements domain.UserDelivery.
func (u *userDelivery) Create(c *fiber.Ctx) (err error) {
	var userData model.UserCreate
	if err := c.BodyParser(&userData); err != nil {
		log.Error().Msgf("")
		return err
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()
	err = u.uc.Create(ctx, userData)

	return
}

// Delete implements domain.UserDelivery.
func (u *userDelivery) Delete(c *fiber.Ctx) (err error) {
	id := c.Params("id")
	userID := "lorem"

	var params = model.UserParams{ID: id, UserID: userID}
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	err = u.uc.Delete(ctx, params)

	return
}

// Find implements domain.UserDelivery.
func (u *userDelivery) Find(c *fiber.Ctx) (err error) {
	return
}

// FindOne implements domain.UserDelivery.
func (u *userDelivery) FindOne(c *fiber.Ctx) (err error) {
	return
}

// Update implements domain.UserDelivery.
func (u *userDelivery) Update(c *fiber.Ctx) (err error) {
	return
}
