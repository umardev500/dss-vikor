package repository

import (
	"context"

	"github.com/umardev500/spk/config"
	"github.com/umardev500/spk/domain"
	"github.com/umardev500/spk/domain/model"
)

type alternateRepository struct {
	trx *config.Trx
}

func NewAlternateRepository(trx *config.Trx) domain.AlternateRepository {
	return &alternateRepository{
		trx: trx,
	}
}

func (u *alternateRepository) Create(ctx context.Context, altrnt model.AlternateCreate) error {
	query := `--sql
		INSERT INTO alternates (id, name, role_id, str, experience, dob, address)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	db := u.trx.GetConn(ctx)
	_, err := db.ExecContext(
		ctx,
		query,
		altrnt.ID,
		altrnt.Name,
		altrnt.RoleID,
		altrnt.STR,
		altrnt.Experience,
		altrnt.DOB,
		altrnt.Address,
	)

	return err
}
