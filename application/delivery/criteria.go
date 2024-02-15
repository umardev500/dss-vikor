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

type criteriaDelivery struct {
	uc domain.CriteriaUsecase
	v  *validator.Validate
}

func NewCriteriaDelivery(uc domain.CriteriaUsecase, r fiber.Router, v *validator.Validate) {
	handler := &criteriaDelivery{
		uc: uc,
		v:  v,
	}

	r.Post("/", handler.Create)
	r.Delete("/:id", handler.Delete)
	r.Get("/", handler.Find)
	r.Get("/:id", handler.FindById)
	r.Put("/:id", handler.Update)
}

func (c *criteriaDelivery) Create(ctx *fiber.Ctx) error {
	var criteria model.CriteriaCreate
	if err := ctx.BodyParser(&criteria); err != nil {
		uid := uuid.New()
		resp := model.Response{
			ID:      uid,
			Status:  fiber.StatusBadRequest,
			Success: false,
			Message: err.Error(),
		}
		bodyRaw, hndl, err := utils.GetRawBodySingleLine(ctx)
		if err != nil {
			return hndl
		}
		logData := utils.LogBuilder(uid, "failed to parse request body", bodyRaw, err)
		log.Error().Msg(logData)
		return ctx.JSON(resp)
	}

	fields, err := utils.ValidateStruct(c.v, criteria)
	if err != nil {
		uid := uuid.New()
		resp := model.Response{
			ID:      uid,
			Status:  fiber.StatusUnprocessableEntity,
			Success: false,
			Message: "validation error",
			Fields:  fields,
		}
		bodyRaw, hndl, err := utils.GetRawBodySingleLine(ctx)
		if err != nil {
			return hndl
		}

		logData := utils.LogBuilder(uid, fiber.ErrUnprocessableEntity.Message, bodyRaw, err)
		log.Error().Msg(logData)
		return ctx.JSON(resp)
	}

	ctxF, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	resp := c.uc.Create(ctxF, criteria)
	return ctx.JSON(resp)
}

func (c *criteriaDelivery) Delete(ctx *fiber.Ctx) error {
	id, hndl := utils.ParseUUID(ctx)
	if id == nil {
		return hndl
	}

	ctxF, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	resp := c.uc.Delete(ctxF, *id)
	return ctx.JSON(resp)
}

func (c *criteriaDelivery) Find(ctx *fiber.Ctx) error {
	ctxF, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	find := model.CriteriaFind{}
	resp := c.uc.Find(ctxF, find)
	return ctx.JSON(resp)
}

func (c *criteriaDelivery) FindById(ctx *fiber.Ctx) error {
	id, hndl := utils.ParseUUID(ctx)
	if id == nil {
		return hndl
	}

	ctxF, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	resp := c.uc.FindById(ctxF, *id)
	return ctx.JSON(resp)
}

func (c *criteriaDelivery) Update(ctx *fiber.Ctx) error {
	id, hndl := utils.ParseUUID(ctx)
	if id == nil {
		return hndl
	}

	var payload model.CriteriaUpdate
	if err := ctx.BodyParser(&payload); err != nil {
		uid := uuid.New()
		resp := model.Response{
			ID:      uid,
			Status:  fiber.StatusBadRequest,
			Success: false,
			Message: fiber.ErrBadRequest.Message,
		}
		bodyRaw, hndl, err := utils.GetRawBodySingleLine(ctx)
		if err != nil {
			return hndl
		}
		logData := utils.LogBuilder(uid, "failed to parse request body", bodyRaw, err)
		log.Error().Msg(logData)
		return ctx.JSON(resp)
	}
	payload.ID = *id

	ctxF, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	resp := c.uc.Update(ctxF, payload)
	return ctx.JSON(resp)
}
