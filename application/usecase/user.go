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

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsercase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) Create(ctx context.Context, user model.UserCreate) model.Response {
	var response model.Response
	var uid uuid.UUID = uuid.New()
	user.Data.ID = uid
	var pass, err = utils.GeneratBcryptHash(user.Data.Password)
	if err != nil {
		userMsg := "failed to create user"
		response = utils.ResponseBuilder(uid, fiber.StatusInternalServerError, false, userMsg, nil)
		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(user), err)
		log.Error().Msg(msg)
		return response
	}

	user.Data.Password = pass

	statusIsEmpty := user.Data.Status == ""
	if statusIsEmpty {
		user.Data.Status = constants.Inactive
	}

	err = u.repo.Create(ctx, user)
	if err != nil {
		userMsg := "failed to create user"
		pqErr := utils.ParsePostgresError(err)
		utils.CombinePqErr(pqErr, &userMsg)
		response = utils.ResponseBuilder(uid, fiber.StatusInternalServerError, false, userMsg, nil)

		msg := utils.LogBuilder(uid, userMsg, utils.StructToJson(user), err)
		log.Error().Msg(msg)
		return response
	}

	response = utils.ResponseBuilder(uid, fiber.StatusCreated, true, "user created successfully", nil)

	return response
}

func (u *userUsecase) Delete(ctx context.Context, params model.UserParams) model.Response {
	var response model.Response

	err := u.repo.Delete(ctx, params)
	uid := uuid.New()
	if err != nil {
		userMsg := "failed to delete user"

		// Handle not affected error
		if err == constants.ErrorNotAffected {
			userMsg = "failed to delete user, user not found"
			response = utils.ResponseBuilder(uid, fiber.StatusNotFound, false, userMsg, nil)
		} else {
			pqErr := utils.ParsePostgresError(err)
			utils.CombinePqErr(pqErr, &userMsg)
			response = utils.ResponseBuilder(uid, fiber.StatusInternalServerError, false, userMsg, nil)
		}

		logData := utils.LogBuilder(uid, userMsg, utils.StructToJson(params), err)
		log.Error().Msg(logData)
		return response
	}

	response = utils.ResponseBuilder(uid, fiber.StatusCreated, true, "user deleted successfully", nil)

	return response
}

func (u *userUsecase) Find(ctx context.Context, params model.UserFind) (res model.Response) {
	users, err := u.repo.Find(ctx, params)
	uid := uuid.New()
	if err != nil {
		userMsg := "failed to find users"

		pqErr := utils.ParsePostgresError(err)
		utils.CombinePqErr(pqErr, &userMsg)

		res = utils.ResponseBuilder(uid, fiber.StatusInternalServerError, false, userMsg, nil)
		logData := utils.LogBuilder(uid, userMsg, utils.StructToJson(params), err)
		log.Error().Msg(logData)
		return res
	}

	res = utils.ResponseBuilder(uid, fiber.StatusOK, true, "users found successfully", users)
	return
}

func (u *userUsecase) FindOne(ctx context.Context, params model.UserFind) (res model.Response) {
	user, err := u.repo.FindOne(ctx, params)
	uid := uuid.New()
	if err != nil {
		// Handle error no data
		if err == sql.ErrNoRows {
			userMsg := "failed to find user, data not found"
			res = utils.ResponseBuilder(uid, fiber.StatusNotFound, false, userMsg, nil)
			logData := utils.LogBuilder(uid, userMsg, utils.StructToJson(params), err)
			log.Error().Msg(logData)
			return res
		}

		userMsg := "failed to find user"
		pqErr := utils.ParsePostgresError(err)
		utils.CombinePqErr(pqErr, &userMsg)

		res = utils.ResponseBuilder(uid, fiber.StatusInternalServerError, false, userMsg, nil)
		logData := utils.LogBuilder(uid, userMsg, utils.StructToJson(params), err)
		log.Error().Msg(logData)
		return res
	}

	res = utils.ResponseBuilder(uid, fiber.StatusOK, true, "user found successfully", user)

	return
}

func (u *userUsecase) Update(ctx context.Context, params model.UserUpdate) (res model.Response) {
	err := u.repo.Update(ctx, params)
	if err != nil {
		uid := uuid.New()

		userMsg := "failed to update user"
		if err == constants.ErrorNotAffected {
			userMsg = "failed to update user, user not found"
			res = utils.ResponseBuilder(uid, fiber.StatusNotFound, false, userMsg, nil)
		} else {
			pqErr := utils.ParsePostgresError(err)
			utils.CombinePqErr(pqErr, &userMsg)
			res = utils.ResponseBuilder(uid, fiber.StatusInternalServerError, false, userMsg, nil)
		}

		logData := utils.LogBuilder(uid, userMsg, utils.StructToJson(params), err)
		log.Error().Msg(logData)
		return res

	}

	res = utils.ResponseBuilder(uuid.New(), fiber.StatusCreated, true, "user updated successfully", nil)

	return
}
