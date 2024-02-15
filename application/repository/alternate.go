package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/umardev500/spk/config"
	"github.com/umardev500/spk/constants"
	"github.com/umardev500/spk/domain"
	"github.com/umardev500/spk/domain/model"
	"github.com/umardev500/spk/utils"
)

type alternateRepository struct {
	trx *config.Trx
}

func NewAlternateRepository(trx *config.Trx) domain.AlternateRepository {
	return &alternateRepository{
		trx: trx,
	}
}

// Create creates a new alternate record in the database.
//
// Parameters:
//
//	ctx context.Context - The context for the operation.
//	altrnt model.AlternateCreate - The alternate model for creating a new record.
//
// Returns:
//
//	error - An error, if any, during the operation.
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

// Delete removes an alternate record from the database by its ID.
//
// Parameters:
//
//	ctx context.Context - The context for the operation.
//	id uuid.UUID - The unique identifier of the alternate record to be deleted.
//
// Returns:
//
//	error - An error, if any, during the operation.
func (u *alternateRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `--sql
		DELETE FROM alternates WHERE id = $1
	`
	db := u.trx.GetConn(ctx)
	result, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return constants.ErrorNotAffected
	}

	return nil
}

// Find retrieves a list of alternates from the database based on the given filter criteria.
//
// Parameters:
//
//	ctx context.Context - The context for the operation.
//	find model.AlternateFind - The filter criteria used to find alternates.
//
// Returns:
//
//	[]model.Alternate - A slice of Alternate models that meet the filter criteria.
//	error - An error, if any, during the operation.
func (u *alternateRepository) Find(ctx context.Context, find model.AlternateFind) ([]model.Alternate, error) {
	query := `--sql
		SELECT * FROM alternates WHERE 1=1
	`

	// Build query based on filter criteria
	args := []interface{}{}

	db := u.trx.GetConn(ctx)
	rows, err := db.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alternates []model.Alternate
	for rows.Next() {
		var altrnt model.Alternate
		if err := rows.StructScan(&altrnt); err != nil {
			return nil, err
		}
		alternates = append(alternates, altrnt)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return alternates, nil
}

// FindById retrieves an alternate from the database by its ID.
// It takes a context.Context and the ID of the alternate to retrieve.
// It returns the retrieved alternate and any error encountered.
func (a *alternateRepository) FindById(ctx context.Context, id uuid.UUID) (altrnt model.Alternate, err error) {
	query := `--sql
		SELECT * FROM alternates WHERE id = $1
	`
	db := a.trx.GetConn(ctx)
	err = db.QueryRowxContext(ctx, query, id).StructScan(&altrnt)
	return
}

func (a *alternateRepository) Update(ctx context.Context, update model.AlternateUpdate) (err error) {
	var info utils.ColumnArgs
	info.Query = `--sql
		UPDATE alternates SET
	`
	info.Args = append(info.Args, update.ID)
	utils.ParseColumn(update, 1, &info)
	info.Query += ` WHERE id = $1;`

	db := a.trx.GetConn(ctx)
	result, err := db.ExecContext(ctx, info.Query, info.Args...)
	if err != nil {
		return
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return constants.ErrorNotAffected
	}

	return
}
