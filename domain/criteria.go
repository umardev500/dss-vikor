package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/spk/domain/model"
)

type CriteriaUsecase interface {
	Create(ctx context.Context, criteria model.CriteriaCreate) model.Response
	Delete(ctx context.Context, id uuid.UUID) model.Response
	Find(ctx context.Context, find model.CriteriaFind) model.Response
	FindById(ctx context.Context, id uuid.UUID) model.Response
	Update(ctx context.Context, criteria model.CriteriaUpdate) model.Response
}

type CriteriaRepository interface {
	Create(ctx context.Context, criteria model.CriteriaCreate) error
	Delete(ctx context.Context, id uuid.UUID) error
	Find(ctx context.Context, find model.CriteriaFind) ([]model.Criteria, error)
	FindById(ctx context.Context, id uuid.UUID) (model.Criteria, error)
	Update(ctx context.Context, criteria model.CriteriaUpdate) error
}
