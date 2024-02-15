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

type criteriaUsecase struct {
	repo domain.CriteriaRepository
}

func NewCriteriaUsecase(repo domain.CriteriaRepository) domain.CriteriaUsecase {
	return &criteriaUsecase{
		repo: repo,
	}
}

func (u *criteriaUsecase) Create(ctx context.Context, criteria model.CriteriaCreate) (resp model.Response) {
	uid := uuid.New()
	criteria.ID = uid

	err := u.repo.Create(ctx, criteria)
	if err != nil {
		userMsg := "failed to create criteria"
		pqErr := utils.ParsePostgresError(err)
		utils.CombinePqErr(pqErr, &userMsg)
		resp = model.Response{
			ID:      uid,
			Status:  fiber.StatusInternalServerError,
			Success: false,
			Message: userMsg,
		}
		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(criteria), err)
		log.Error().Msg(msg)
		return
	}

	resp = model.Response{
		ID:      uid,
		Status:  fiber.StatusInternalServerError,
		Success: false,
		Message: "criteria created successfully",
	}

	return resp
}

func (a *criteriaUsecase) Delete(ctx context.Context, id uuid.UUID) (resp model.Response) {
	err := a.repo.Delete(ctx, id)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to delete criteria"
		if err == constants.ErrorNotAffected {
			userMsg = "failed to delete criteria, criteria not found"
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
		Message: "criteria deleted successfully",
	}
}

func (a *criteriaUsecase) Find(ctx context.Context, find model.CriteriaFind) (resp model.Response) {
	criterias, err := a.repo.Find(ctx, find)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to find criterias"

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
		Message: "find criterias",
		Data:    criterias,
	}
}

func (a *criteriaUsecase) FindById(ctx context.Context, id uuid.UUID) (resp model.Response) {
	criteria, err := a.repo.FindById(ctx, id)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to find criteria"

		if err == sql.ErrNoRows {
			userMsg = "failed to find criteria, criteria not found"
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
		Message: "find criteria",
		Data:    criteria,
	}
}

func (a *criteriaUsecase) Update(ctx context.Context, criteria model.CriteriaUpdate) (resp model.Response) {
	err := a.repo.Update(ctx, criteria)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to update criteria"

		if err == constants.ErrorNotAffected {
			userMsg = "failed to update criteria, criteria not found"
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

		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(criteria), err)
		log.Error().Msg(msg)
		return
	}

	resp = model.Response{
		ID:      uuid.New(),
		Status:  fiber.StatusCreated,
		Success: true,
		Message: "criteria updated successfully",
	}

	return
}
