package usecase

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/spk/domain"
	"github.com/umardev500/spk/domain/model"
	"github.com/umardev500/spk/utils"
)

type roleUsecase struct {
	repo domain.RoleRepository
}

func NewRoleUsecase(repo domain.RoleRepository) domain.RoleUsecase {
	return &roleUsecase{
		repo: repo,
	}
}

func (u *roleUsecase) Create(ctx context.Context, role model.RoleCreate) model.Response {
	var response model.Response

	uid := uuid.New()
	role.ID = uid

	err := u.repo.Create(ctx, role)
	if err != nil {
		userMsg := "failed to create role"
		pqErr := utils.ParsePostgresError(err)
		if pqErr != nil {
			utils.CombinePqErr(pqErr.Error(), &userMsg)
		}
		response = utils.ResponseBuilder(uid, fiber.StatusInternalServerError, false, userMsg, nil)
		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(role), err)
		log.Error().Msg(msg)
		return response
	}

	return utils.ResponseBuilder(uid, fiber.StatusCreated, true, "role created successfully", nil)
}
