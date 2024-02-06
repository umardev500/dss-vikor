package usecase

import (
	"context"
	"fmt"

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
	var uuid uuid.UUID = uuid.New()
	user.Data.ID = uuid

	statusIsEmpty := user.Data.Status == ""
	if statusIsEmpty {
		user.Data.Status = constants.Inactive
	}

	err := u.repo.Create(ctx, user)
	if err != nil {
		userMsg := "failed to create user"
		pqErr := utils.ParsePostgresError(err)
		if pqErr != nil {
			utils.CombinePqErr(pqErr.Error(), &userMsg)
		}
		response = utils.ResponseBuilder(uuid, fiber.StatusInternalServerError, false, userMsg, nil)

		msg := utils.LogBuilder(uuid, userMsg, utils.StructToJson(user), err)
		log.Error().Msg(msg)
		return response
	}

	response = utils.ResponseBuilder(uuid, fiber.StatusCreated, true, "user created successfully", nil)

	return response
}

func (u *userUsecase) Delete(ctx context.Context, params model.UserParams) model.Response {
	var response model.Response

	err := u.repo.Delete(ctx, params)
	if err != nil {
		uuid := uuid.New()
		userMsg := "failed to delete user"

		// Handle not affected error
		if err == constants.ErrorNotAffected {
			userMsg = "failed to delete user, user not found"
			response = utils.ResponseBuilder(uuid, fiber.StatusNotFound, false, userMsg, nil)
			logData := utils.LogBuilder(uuid, userMsg, utils.StructToJson(params), err)
			log.Error().Msg(logData)
			return response
		}

		pqErr := utils.ParsePostgresError(err)
		if pqErr != nil {
			utils.CombinePqErr(pqErr.Error(), &userMsg)
		}
		response = utils.ResponseBuilder(uuid, fiber.StatusInternalServerError, false, userMsg, nil)
		logData := utils.LogBuilder(uuid, userMsg, utils.StructToJson(params), err)
		log.Error().Msg(logData)
		return response
	}

	response = utils.ResponseBuilder(params.UserID, fiber.StatusCreated, true, "user deleted successfully", nil)

	return response
}

func (u *userUsecase) Find(ctx context.Context, params model.UserFind) (res model.Response) {
	users, err := u.repo.Find(ctx, params)
	if err != nil {
		uuid := uuid.New()
		userMsg := "failed to find users"

		pqErr := utils.ParsePostgresError(err)
		if pqErr != nil {
			utils.CombinePqErr(pqErr.Error(), &userMsg)
		}

		res = utils.ResponseBuilder(uuid, fiber.StatusInternalServerError, false, userMsg, nil)
		logData := utils.LogBuilder(uuid, userMsg, utils.StructToJson(params), err)
		log.Error().Msg(logData)
		return res
	}

	res = utils.ResponseBuilder(uuid.New(), fiber.StatusOK, true, "users found successfully", users)
	return
}

func (u *userUsecase) FindOne(ctx context.Context, params model.UserFind) (res model.Response, err error) {
	user, err := u.repo.FindOne(ctx, params)
	fmt.Println(user)

	return
}

func (u *userUsecase) Update(ctx context.Context, params model.UserUpdate) (err error) {
	err = u.repo.Update(ctx, params)

	return
}
