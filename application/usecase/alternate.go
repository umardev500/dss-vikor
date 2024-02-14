package usecase

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/spk/domain"
	"github.com/umardev500/spk/domain/model"
	"github.com/umardev500/spk/utils"
)

type alternateUsecase struct {
	repo domain.AlternateRepository
}

func NewAlternateUsecase(repo domain.AlternateRepository) domain.AlternateUsecase {
	return &alternateUsecase{
		repo: repo,
	}
}

func (u *alternateUsecase) Create(ctx context.Context, altrnt model.AlternateCreate) (resp model.Response) {
	uid := uuid.New()
	altrnt.ID = uid

	fmt.Println(altrnt)

	err := u.repo.Create(ctx, altrnt)
	if err != nil {
		userMsg := "failed to create alternate"
		pqErr := utils.ParsePostgresError(err)
		utils.CombinePqErr(pqErr, &userMsg)
		resp = utils.ResponseBuilder(uid, fiber.StatusInternalServerError, false, userMsg, nil)
		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(altrnt), err)
		log.Error().Msg(msg)
		return
	}

	return utils.ResponseBuilder(uid, fiber.StatusCreated, true, "alternate created successfully", nil)
}
