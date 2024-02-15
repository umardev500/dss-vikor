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

type subCriteriaUsecase struct {
	repo domain.SubCriteriaRepository
}

func NewSubCriteriaUsecase(repo domain.SubCriteriaRepository) domain.SubCriteriaUsecase {
	return &subCriteriaUsecase{
		repo: repo,
	}
}

func (u *subCriteriaUsecase) Create(ctx context.Context, subCriteria model.SubCriteriaCreate) (resp model.Response) {
	uid := uuid.New()
	subCriteria.ID = uid

	err := u.repo.Create(ctx, subCriteria)
	if err != nil {
		userMsg := "failed to create sub-criteria"
		pqErr := utils.ParsePostgresError(err)
		utils.CombinePqErr(pqErr, &userMsg)
		resp = model.Response{
			ID:      uid,
			Status:  fiber.StatusInternalServerError,
			Success: false,
			Message: userMsg,
		}
		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(subCriteria), err)
		log.Error().Msg(msg)
		return
	}

	resp = model.Response{
		ID:      uid,
		Status:  fiber.StatusInternalServerError,
		Success: false,
		Message: "sub-criteria created successfully",
	}

	return resp
}

func (u *subCriteriaUsecase) Delete(ctx context.Context, id uuid.UUID) (resp model.Response) {
	err := u.repo.Delete(ctx, id)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to delete sub-criteria"
		if err == constants.ErrorNotAffected {
			userMsg = "failed to delete sub-criteria, sub-criteria not found"
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
		Message: "sub-criteria deleted successfully",
	}
}

func (u *subCriteriaUsecase) Find(ctx context.Context, find model.SubCriteriaFind) (resp model.Response) {
	subCriteria, err := u.repo.Find(ctx, find)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to find sub-criteria"

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
		Message: "find sub-criteria",
		Data:    subCriteria,
	}
}

func (u *subCriteriaUsecase) FindById(ctx context.Context, id uuid.UUID) (resp model.Response) {
	subCriteria, err := u.repo.FindById(ctx, id)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to find sub-criteria"

		if err == sql.ErrNoRows {
			userMsg = "failed to find sub-criteria, sub-criteria not found"
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
		Message: "find sub-criteria",
		Data:    subCriteria,
	}
}

func (u *subCriteriaUsecase) Update(ctx context.Context, subCriteria model.SubCriteriaUpdate) (resp model.Response) {
	err := u.repo.Update(ctx, subCriteria)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to update sub-criteria"

		if err == constants.ErrorNotAffected {
			userMsg = "failed to update sub-criteria, sub-criteria not found"
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

		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(subCriteria), err)
		log.Error().Msg(msg)
		return
	}

	resp = model.Response{
		ID:      uuid.New(),
		Status:  fiber.StatusCreated,
		Success: true,
		Message: "sub-criteria updated successfully",
	}

	return
}
