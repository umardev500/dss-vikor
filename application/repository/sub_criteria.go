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

type subCriteriaRepository struct {
	trx *config.Trx
}

func NewSubCriteriaRepository(trx *config.Trx) domain.SubCriteriaRepository {
	return &subCriteriaRepository{
		trx: trx,
	}
}

func (r *subCriteriaRepository) Create(ctx context.Context, subCriteria model.SubCriteriaCreate) error {
	query := `--sql
        INSERT INTO sub_criterias (id, criteria_id, name) VALUES ($1, $2, $3);
    `

	db := r.trx.GetConn(ctx)
	_, err := db.ExecContext(ctx, query, subCriteria.ID, subCriteria.CriteriaID, subCriteria.Name)
	return err
}

func (r *subCriteriaRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `--sql
        DELETE FROM sub_criterias WHERE id = $1;
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

func (r *subCriteriaRepository) Find(ctx context.Context, find model.SubCriteriaFind) ([]model.SubCriteria, error) {
	query := `--sql
        SELECT * FROM sub_criterias WHERE 1=1;
    `

	db := r.trx.GetConn(ctx)
	rows, err := db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subCriteriaList = make([]model.SubCriteria, 0)
	for rows.Next() {
		var subCriteria model.SubCriteria
		if err := rows.StructScan(&subCriteria); err != nil {
			return nil, err
		}
		subCriteriaList = append(subCriteriaList, subCriteria)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return subCriteriaList, nil
}

func (r *subCriteriaRepository) FindById(ctx context.Context, id uuid.UUID) (model.SubCriteria, error) {
	var subCriteria model.SubCriteria
	query := `--sql
        SELECT * FROM sub_criterias WHERE id = $1;
    `

	db := r.trx.GetConn(ctx)
	err := db.QueryRowxContext(ctx, query, id).StructScan(&subCriteria)
	if err != nil {
		return model.SubCriteria{}, err
	}
	return subCriteria, nil
}

func (r *subCriteriaRepository) Update(ctx context.Context, subCriteria model.SubCriteriaUpdate) error {
	var info utils.ColumnArgs
	info.Query = `--sql
        UPDATE sub_criterias SET
    `
	info.Args = append(info.Args, subCriteria.ID)
	utils.ParseColumn(subCriteria, 1, &info)
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
