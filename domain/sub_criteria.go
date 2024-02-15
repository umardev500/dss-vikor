package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/spk/domain/model"
)

type SubCriteriaUsecase interface {
	Create(ctx context.Context, subCriteria model.SubCriteriaCreate) model.Response
	Delete(ctx context.Context, id uuid.UUID) model.Response
	Find(ctx context.Context, find model.SubCriteriaFind) model.Response
	FindById(ctx context.Context, id uuid.UUID) model.Response
	Update(ctx context.Context, subCriteria model.SubCriteriaUpdate) model.Response
}

type SubCriteriaRepository interface {
	Create(ctx context.Context, subCriteria model.SubCriteriaCreate) error
	Delete(ctx context.Context, id uuid.UUID) error
	Find(ctx context.Context, find model.SubCriteriaFind) ([]model.SubCriteria, error)
	FindById(ctx context.Context, id uuid.UUID) (model.SubCriteria, error)
	Update(ctx context.Context, subCriteria model.SubCriteriaUpdate) error
}
