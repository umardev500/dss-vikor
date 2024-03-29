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

type roleUsecase struct {
	repo domain.RoleRepository
}

func NewRoleUsecase(repo domain.RoleRepository) domain.RoleUsecase {
	return &roleUsecase{
		repo: repo,
	}
}

// Create creates a new role.
//
// ctx context.Context, role model.RoleCreate
// model.Response
func (u *roleUsecase) Create(ctx context.Context, role model.RoleCreate) model.Response {
	var response model.Response

	uid := uuid.New()
	role.ID = uid

	err := u.repo.Create(ctx, role)
	if err != nil {
		userMsg := "failed to create role"
		pqErr := utils.ParsePostgresError(err)
		utils.CombinePqErr(pqErr, &userMsg)
		response = model.Response{
			ID:      uid,
			Status:  fiber.StatusInternalServerError,
			Success: false,
			Message: userMsg,
		}
		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(role), err)
		log.Error().Msg(msg)
		return response
	}

	return utils.ResponseBuilder(uid, fiber.StatusCreated, true, "role created successfully", nil)
}

// Delete deletes a role using the given ID.
//
// ctx context.Context, id uuid.UUID
// model.Response
func (u *roleUsecase) Delete(ctx context.Context, id uuid.UUID) model.Response {
	var response model.Response

	err := u.repo.Delete(ctx, id)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to delete role"

		if err == constants.ErrorNotAffected {
			userMsg = "failed to delete role, role not found"
			response = model.Response{
				ID:      uid,
				Status:  fiber.StatusNotFound,
				Success: false,
				Message: userMsg,
			}
		} else {
			pqErr := utils.ParsePostgresError(err)
			utils.CombinePqErr(pqErr, &userMsg)
			response = model.Response{
				ID:      uid,
				Status:  fiber.StatusInternalServerError,
				Success: false,
				Message: userMsg,
			}
		}

		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(id), err)
		log.Error().Msg(msg)
		return response
	}

	return utils.ResponseBuilder(id, fiber.StatusCreated, true, "role deleted successfully", nil)
}

// Find is a function that retrieves roles based on the given criteria.
//
// It takes a context and a model.RoleFind as parameters and returns a model.Response.
func (u *roleUsecase) Find(ctx context.Context, find model.RoleFind) (resp model.Response) {
	roles, err := u.repo.Find(ctx, find)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to find roles"

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

	resp = model.Response{
		ID:      uuid.New(),
		Status:  fiber.StatusOK,
		Success: true,
		Message: "find roles",
		Data:    roles,
	}

	return
}

// FindById finds a role by its ID.
//
// ctx context.Context, id uuid.UUID
// resp model.Response
func (u *roleUsecase) FindById(ctx context.Context, id uuid.UUID) (resp model.Response) {
	role, err := u.repo.FindById(ctx, id)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to find role"

		if err == sql.ErrNoRows {
			userMsg = "failed to find roles, roles not found"
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

	resp = utils.ResponseBuilder(uuid.New(), fiber.StatusOK, true, "find role", role)
	return
}

// Update updates a role in the database.
//
// ctx context.Context, id uuid.UUID, role model.RoleUpdate
// resp model.Response
func (r *roleUsecase) Update(ctx context.Context, id uuid.UUID, role model.RoleUpdate) (resp model.Response) {
	err := r.repo.Update(ctx, id, role)
	if err != nil {
		uid := uuid.New()
		userMsg := "failed to update role"

		if err == constants.ErrorNotAffected {
			userMsg = "failed to update role, role not found"
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

	return utils.ResponseBuilder(uuid.New(), fiber.StatusCreated, true, "role updated successfully", nil)
}
