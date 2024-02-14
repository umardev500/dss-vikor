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

	err := u.repo.Create(ctx, altrnt)
	if err != nil {
		userMsg := "failed to create alternate"
		pqErr := utils.ParsePostgresError(err)
		utils.CombinePqErr(pqErr, &userMsg)
		resp = model.Response{
			ID:      uid,
			Status:  fiber.StatusInternalServerError,
			Success: false,
			Message: userMsg,
		}
		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(altrnt), err)
		log.Error().Msg(msg)
		return
	}

	resp = model.Response{
		ID:      uid,
		Status:  fiber.StatusInternalServerError,
		Success: false,
		Message: "alternate created successfully",
	}

	return resp
}
