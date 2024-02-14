package domain

import (
	"context"

	"github.com/umardev500/spk/domain/model"
)

type AlternateUsecase interface {
	Create(ctx context.Context, altrnt model.AlternateCreate) model.Response
}

type AlternateRepository interface {
	Create(ctx context.Context, altrnt model.AlternateCreate) error
}
