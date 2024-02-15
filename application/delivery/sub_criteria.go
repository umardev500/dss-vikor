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

type subCriteriaDelivery struct {
	uc domain.SubCriteriaUsecase
	v  *validator.Validate
}

func NewSubCriteriaDelivery(uc domain.SubCriteriaUsecase, r fiber.Router, v *validator.Validate) {
	handler := &subCriteriaDelivery{
		uc: uc,
		v:  v,
	}

	r.Post("/", handler.Create)
	r.Delete("/:id", handler.Delete)
	r.Get("/", handler.Find)
	r.Get("/:id", handler.FindById)
	r.Put("/:id", handler.Update)
}

func (s *subCriteriaDelivery) Create(c *fiber.Ctx) error {
	var subCriteria model.SubCriteriaCreate
	if err := c.BodyParser(&subCriteria); err != nil {
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

	fields, err := utils.ValidateStruct(s.v, subCriteria)
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

	resp := s.uc.Create(ctx, subCriteria)
	return c.JSON(resp)
}

func (s *subCriteriaDelivery) Delete(c *fiber.Ctx) error {
	id, hndl := utils.ParseUUID(c)
	if id == nil {
		return hndl
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	resp := s.uc.Delete(ctx, *id)
	return c.JSON(resp)
}

func (s *subCriteriaDelivery) Find(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	find := model.SubCriteriaFind{}
	resp := s.uc.Find(ctx, find)
	return c.JSON(resp)
}

func (s *subCriteriaDelivery) FindById(c *fiber.Ctx) error {
	id, hndl := utils.ParseUUID(c)
	if id == nil {
		return hndl
	}

	ctx, cancel := context.WithTimeout(c.Context(), 10*time.Second)
	defer cancel()

	resp := s.uc.FindById(ctx, *id)
	return c.JSON(resp)
}

func (s *subCriteriaDelivery) Update(c *fiber.Ctx) error {
	id, hndl := utils.ParseUUID(c)
	if id == nil {
		return hndl
	}

	var payload model.SubCriteriaUpdate
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

	resp := s.uc.Update(ctx, payload)
	return c.JSON(resp)
}
