package usecase

import (
	"context"
	"fmt"

	"github.com/umardev500/spk/domain"
	"github.com/umardev500/spk/domain/model"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsercase(repo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) Create(ctx context.Context, user model.User) (err error) {
	err = u.repo.Create(ctx, user)

	return
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
