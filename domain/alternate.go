package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/spk/domain/model"
)

type AlternateUsecase interface {
	Create(ctx context.Context, altrnt model.AlternateCreate) model.Response
	Delete(ctx context.Context, id uuid.UUID) model.Response
	Find(ctx context.Context, find model.AlternateFind) model.Response
	FindById(ctx context.Context, id uuid.UUID) model.Response
	Update(ctx context.Context, altrnt model.AlternateUpdate) model.Response
}

type AlternateRepository interface {
	Create(ctx context.Context, altrnt model.AlternateCreate) error
	Delete(ctx context.Context, id uuid.UUID) error
	Find(ctx context.Context, find *model.AlternateFind) ([]model.Alternate, error)
	FindById(ctx context.Context, id uuid.UUID) (model.Alternate, error)
	Update(ctx context.Context, altrnt model.AlternateUpdate) error
}
