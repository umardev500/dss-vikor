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

		msg := fmt.Sprintf("error creating user: %v uuid: %s", err, uuid)
		log.Error().Msg(msg)
	}

	return response
}

func (u *userUsecase) Delete(ctx context.Context, params model.UserParams) (err error) {
	err = u.repo.Delete(ctx, params)

	return
}

func (u *userUsecase) Find(ctx context.Context, params model.UserFind) (res model.Response, err error) {
	users, err := u.repo.Find(ctx, params)
	fmt.Println(users)

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
