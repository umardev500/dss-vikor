package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/spk/config"
	"github.com/umardev500/spk/constants"
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

// Create creates a new role in the role repository.
//
// ctx context.Context, role model.RoleCreate
// error
func (r *roleRepository) Create(ctx context.Context, role model.RoleCreate) error {
	query := `--sql
		INSERT INTO roles (id, name)
		VALUES ($1, $2);
	`

	db := r.trx.GetConn(ctx)

	_, err := db.ExecContext(ctx, query, role.ID, role.Name)
	return err
}

func (r *roleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `--sql
		DELETE FROM roles WHERE id = $1;
	`

	db := r.trx.GetConn(ctx)

	result, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return constants.ErrorNotAffected
	}

	return err
}
