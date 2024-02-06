package delivery

import (
	"context"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/spk/domain"
	"github.com/umardev500/spk/domain/model"
	"github.com/umardev500/spk/utils"
)

type userDelivery struct {
	uc     domain.UserUsecase
	router fiber.Router
}

func NewUserDelivery(uc domain.UserUsecase, router fiber.Router) {
	var handler domain.UserDelivery = &userDelivery{
		uc:     uc,
		router: router,
	}

	router.Post("/", handler.Create)
	router.Delete("/:id", handler.Delete)
	router.Get("/", handler.Find)
}

// Create implements domain.UserDelivery.
func (u *userDelivery) Create(c *fiber.Ctx) (err error) {
	var userData model.UserToCreate
	if err := c.BodyParser(&userData); err != nil {
		uuid := uuid.New()
		resp := utils.ResponseBuilder(uuid, fiber.StatusBadRequest, false, err.Error(), nil)
		bodyRaw := string(c.BodyRaw())
		logData := utils.LogBuilder(uuid, "failed to parse request body", bodyRaw, err)
		log.Error().Msg(logData)

		return c.JSON(resp)
	}

	var userParams = model.UserCreate{
		Data: &userData,
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()
	resp := u.uc.Create(ctx, userParams)

	return c.JSON(resp)
}

// Delete implements domain.UserDelivery.
func (u *userDelivery) Delete(c *fiber.Ctx) (err error) {
	id := c.Params("id")
	uid, err := uuid.Parse(id)

	if err != nil {
		debugID := uuid.New()
		resp := utils.ResponseBuilder(debugID, fiber.StatusBadRequest, false, err.Error(), nil)
		bodyRaw := string(c.BodyRaw())
		logData := utils.LogBuilder(debugID, "failed to parse request body", bodyRaw, err)
		log.Error().Msg(logData)

		return c.JSON(resp)
	}

	userID := uuid.New()

	var params = model.UserParams{ID: uid, UserID: userID}
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	resp := u.uc.Delete(ctx, params)

	return c.JSON(resp)
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
