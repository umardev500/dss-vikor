package domain

import (
	"context"

	"github.com/umardev500/spk/domain/model"
)

type UserRepository interface {
	Create(context.Context, model.User) error
	Delete(context.Context, model.UserParams) error
	Find(context.Context, model.UserFind) ([]model.User, error)
	FindOne(context.Context, model.UserFind) (model.User, error)
	Update(context.Context, model.UserUpdate) error
}
