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

// Delete deletes a role from the repository.
//
// ctx - the context
// id - the UUID of the role to delete
// error - returns an error, if any
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

// Find finds roles based on the provided criteria.
//
// ctx context.Context, find model.RoleFind
// []model.Role, error
func (r *roleRepository) Find(ctx context.Context, find model.RoleFind) (roles []model.Role, err error) {
	query := `--sql
		SELECT * FROM roles
	`
	db := r.trx.GetConn(ctx)

	cur, err := db.QueryxContext(ctx, query)
	if err != nil {
		return
	}

	roles = make([]model.Role, 0)

	for cur.Next() {
		var each model.Role
		if err := cur.StructScan(&each); err != nil {
			return nil, err
		}

		roles = append(roles, each)
	}

	return
}

// FindById description of the Go function.
//
// ctx context.Context, id uuid.UUID
// role model.Role, err error
func (r *roleRepository) FindById(ctx context.Context, id uuid.UUID) (role model.Role, err error) {
	query := `--sql
		SELECT * FROM roles WHERE id = $1
	`

	db := r.trx.GetConn(ctx)
	err = db.QueryRowxContext(ctx, query, id).StructScan(&role)

	return
}

// Update updates a role in the role repository.
//
// ctx context.Context, id uuid.UUID, role model.RoleUpdate
// error
func (r *roleRepository) Update(ctx context.Context, id uuid.UUID, role model.RoleUpdate) error {
	query := `--sql
		UPDATE roles
		SET name = $2
		WHERE id = $1
	`

	db := r.trx.GetConn(ctx)
	result, err := db.ExecContext(ctx, query, id, role.Name)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return constants.ErrorNotAffected
	}

	return err
}
