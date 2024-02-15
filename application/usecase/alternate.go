package usecase

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/spk/constants"
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

func (a *alternateUsecase) Delete(ctx context.Context, id uuid.UUID) (resp model.Response) {
	err := a.repo.Delete(ctx, id)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to delete alternate"
		if err == constants.ErrorNotAffected {
			userMsg = "failed to delete alternate, alternate not found"
			resp = model.Response{
				ID:      uid,
				Status:  fiber.StatusNotFound,
				Success: false,
				Message: userMsg,
			}
		} else {
			pqErr := utils.ParsePostgresError(err)
			utils.CombinePqErr(pqErr, &userMsg)
			resp = model.Response{
				ID:      uid,
				Status:  fiber.StatusInternalServerError,
				Success: false,
				Message: userMsg,
			}
		}

		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(id), err)
		log.Error().Msg(msg)
		return
	}

	return model.Response{
		ID:      uuid.New(),
		Status:  fiber.StatusCreated,
		Success: true,
		Message: "alternate deleted successfully",
	}
}

func (a *alternateUsecase) Find(ctx context.Context, find model.AlternateFind) (resp model.Response) {
	alternates, err := a.repo.Find(ctx, find)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to find alternates"

		pqErr := utils.ParsePostgresError(err)
		utils.CombinePqErr(pqErr, &userMsg)
		resp = model.Response{
			ID:      uid,
			Status:  fiber.StatusInternalServerError,
			Success: false,
			Message: userMsg,
		}

		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(find), err)
		log.Error().Msg(msg)
		return
	}

	return model.Response{
		ID:      uuid.New(),
		Status:  fiber.StatusOK,
		Success: true,
		Message: "find alternates",
		Data:    alternates,
	}
}

func (a *alternateUsecase) FindById(ctx context.Context, id uuid.UUID) (resp model.Response) {
	alternates, err := a.repo.FindById(ctx, id)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to find alternate"

		if err == sql.ErrNoRows {
			userMsg = "failed to find alternate, alternate not found"
			resp = model.Response{
				ID:      uid,
				Status:  fiber.StatusNotFound,
				Success: false,
				Message: userMsg,
			}
		} else {
			pqErr := utils.ParsePostgresError(err)
			utils.CombinePqErr(pqErr, &userMsg)
			resp = model.Response{
				ID:      uid,
				Status:  fiber.StatusInternalServerError,
				Success: false,
				Message: userMsg,
			}
		}

		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(id), err)
		log.Error().Msg(msg)
		return
	}

	return model.Response{
		ID:      uuid.New(),
		Status:  fiber.StatusOK,
		Success: true,
		Message: "find alternate",
		Data:    alternates,
	}
}

func (a *alternateUsecase) Update(ctx context.Context, altrnt model.AlternateUpdate) (resp model.Response) {
	err := a.repo.Update(ctx, altrnt)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to update alternate"

		if err == constants.ErrorNotAffected {
			userMsg = "failed to update alternate, alternate not found"
			resp = model.Response{
				ID:      uid,
				Status:  fiber.StatusNotFound,
				Success: false,
				Message: userMsg,
			}
		} else {
			pqErr := utils.ParsePostgresError(err)
			utils.CombinePqErr(pqErr, &userMsg)
			resp = model.Response{
				ID:      uid,
				Status:  fiber.StatusInternalServerError,
				Success: false,
				Message: userMsg,
			}
		}

		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(altrnt), err)
		log.Error().Msg(msg)
		return
	}

	resp = model.Response{
		ID:      uuid.New(),
		Status:  fiber.StatusCreated,
		Success: true,
		Message: "alternate updated successfully",
	}

	return
}
