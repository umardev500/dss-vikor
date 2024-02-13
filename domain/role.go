package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/spk/domain/model"
)

type RoleUsecase interface {
	Create(ctx context.Context, role model.RoleCreate) model.Response
	Delete(ctx context.Context, id uuid.UUID) model.Response
	Find(ctx context.Context, find model.RoleFind) model.Response
	FindById(ctx context.Context, id uuid.UUID) model.Response
}

type RoleRepository interface {
	Create(ctx context.Context, role model.RoleCreate) error
	Delete(ctx context.Context, id uuid.UUID) error
	Find(ctx context.Context, find model.RoleFind) ([]model.Role, error)
	FindById(ctx context.Context, id uuid.UUID) (model.Role, error)
	Update()
}
