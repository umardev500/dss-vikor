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

type criteriaRepository struct {
	trx *config.Trx
}

func NewCriteriaRepository(trx *config.Trx) domain.CriteriaRepository {
	return &criteriaRepository{
		trx: trx,
	}
}

func (r *criteriaRepository) Create(ctx context.Context, criteria model.CriteriaCreate) error {
	query := `--sql
		INSERT INTO criterias (id, name) VALUES ($1, $2);
	`

	db := r.trx.GetConn(ctx)
	_, err := db.ExecContext(ctx, query, criteria.ID, criteria.Name)
	return err
}

func (r *criteriaRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `--sql
		DELETE FROM criterias WHERE id = $1;
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

	return nil
}

func (r *criteriaRepository) Find(ctx context.Context, find model.CriteriaFind) ([]model.Criteria, error) {
	query := `--sql
		SELECT * FROM criterias WHERE 1=1;
	`

	db := r.trx.GetConn(ctx)
	rows, err := db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var criteriaList = make([]model.Criteria, 0)
	for rows.Next() {
		var criteria model.Criteria
		if err := rows.StructScan(&criteria); err != nil {
			return nil, err
		}
		criteriaList = append(criteriaList, criteria)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return criteriaList, nil
}

func (r *criteriaRepository) FindById(ctx context.Context, id uuid.UUID) (model.Criteria, error) {
	var criteria model.Criteria
	query := `--sql
		SELECT * FROM criterias WHERE id = $1;
	`
	db := r.trx.GetConn(ctx)
	err := db.QueryRowxContext(ctx, query, id).StructScan(&criteria)
	if err != nil {
		return model.Criteria{}, err
	}
	return criteria, nil
}

func (r *criteriaRepository) Update(ctx context.Context, criteria model.CriteriaUpdate) error {
	var info utils.ColumnArgs
	info.Query = `--sql
        UPDATE criterias SET
    `
	info.Args = append(info.Args, criteria.ID)
	utils.ParseColumn(criteria, 1, &info)
	info.Query += ` WHERE id = $1;`

	db := r.trx.GetConn(ctx)
	result, err := db.ExecContext(ctx, info.Query, info.Args...)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return constants.ErrorNotAffected
	}

	return nil
}
