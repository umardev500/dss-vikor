package repository

import (
	"context"

	"github.com/umardev500/spk/config"
	"github.com/umardev500/spk/domain"
	"github.com/umardev500/spk/domain/model"
)

type roleRepository struct {
	trx *config.Trx
}

func NewRoleRepository(trx *config.Trx) domain.RoleRepository {
	return &roleRepository{
		trx: trx,
	}
}

func (r *roleRepository) Create(ctx context.Context, role model.RoleCreate) error {
	query := `--sql
		INSERT INTO roles (id, name)
		VALUES ($1, $2);
	`

	db := r.trx.GetConn(ctx)

	_, err := db.ExecContext(ctx, query, role.ID, role.Name)
	return err
}
