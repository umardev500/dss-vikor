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
	r.Delete("/:id", handler.Delete)
	r.Get("/", handler.Find)
	r.Get("/:id", handler.FindById)
	r.Put("/:id", handler.Update)
}

// Create handles the creation of an alternate delivery.
//
// It takes a fiber context as a parameter and returns an error.
func (a *alternateDelivery) Create(c *fiber.Ctx) error {
	var alternate model.AlternateCreate
	if err := c.BodyParser(&alternate); err != nil {
		uid := uuid.New()
		resp := model.Response{
			ID:      uid,
			Status:  fiber.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		}
		bodyRaw, hndl, err := utils.GetRawBodySingleLine(c)
		if err != nil {
			return hndl
		}
		logData := utils.LogBuilder(uid, "failed to parse request body", bodyRaw, err)
		log.Error().Msg(logData)
		return c.JSON(resp)
	}

	// Validate the struct
	fields, err := utils.ValidateStruct(a.v, alternate)
	if err != nil {
		uid := uuid.New()
		resp := model.Response{
			ID:      uid,
			Status:  fiber.StatusUnprocessableEntity,
			Success: false,
			Message: "validation error",
			Fields:  fields,
		}
		bodyRaw, hndl, err := utils.GetRawBodySingleLine(c)
		if err != nil {
			return hndl
		}

		logData := utils.LogBuilder(uid, fiber.ErrUnprocessableEntity.Message, bodyRaw, err)
		log.Error().Msg(logData)
		return c.JSON(resp)
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	resp := a.uc.Create(ctx, alternate)
	return c.JSON(resp)
}

func (a *alternateDelivery) Delete(c *fiber.Ctx) error {
	id, hndl := utils.ParseUUID(c)
	if id == nil {
		return hndl
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	resp := a.uc.Delete(ctx, *id)
	return c.JSON(resp)
}

func (a *alternateDelivery) Find(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	var pageInfo = model.PageInfo{}
	utils.GetPageInfo(c, &pageInfo)

	find := model.AlternateFind{
		PageInfo: pageInfo,
	}
	resp := a.uc.Find(ctx, find)
	return c.JSON(resp)
}

func (a *alternateDelivery) FindById(c *fiber.Ctx) error {
	id, hndl := utils.ParseUUID(c)
	if id == nil {
		return hndl
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	resp := a.uc.FindById(ctx, *id)
	return c.JSON(resp)
}

func (a *alternateDelivery) Update(c *fiber.Ctx) error {
	id, hndl := utils.ParseUUID(c)
	if id == nil {
		return hndl
	}

	var payload model.AlternateUpdate
	if err := c.BodyParser(&payload); err != nil {
		uid := uuid.New()
		resp := model.Response{
			ID:      uid,
			Status:  fiber.StatusBadRequest,
			Success: false,
			Message: fiber.ErrBadRequest.Message,
		}
		bodyRaw, hndl, err := utils.GetRawBodySingleLine(c)
		if err != nil {
			return hndl
		}
		logData := utils.LogBuilder(uid, "failed to parse request body", bodyRaw, err)
		log.Error().Msg(logData)
		return c.JSON(resp)
	}
	payload.ID = *id

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	resp := a.uc.Update(ctx, payload)
	return c.JSON(resp)
}
