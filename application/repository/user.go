package repository

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/spk/config"
	"github.com/umardev500/spk/constants"
	"github.com/umardev500/spk/domain"
	"github.com/umardev500/spk/domain/model"
)

type userRepository struct {
	trx *config.Trx
}

func NewUserRepository(trx *config.Trx) domain.UserRepository {
	return &userRepository{
		trx: trx,
	}
}

func (u *userRepository) Create(ctx context.Context, data model.User) (err error) {
	log.Debug().Msgf("creating user: %v", data)
	query := `--sql
		INSERT INTO users (id, email, password, status)
		VALUES ($1, $2, $2, $4);
	`

	db := u.trx.GetConn(ctx)
	_, err = db.ExecContext(ctx, query, data.ID, data.Email, data.Password, data.Status)
	if err != nil {
		log.Error().Msgf("error creating user: %v", err)
		return
	}

	log.Debug().Msgf("created user: %v", data)

	return
}

func (u *userRepository) Delete(ctx context.Context, params model.UserParams) (err error) {
	log.Debug().Msgf("deleting user")
	query := `--sql
		DELETE FROM users WHERE id = $1;
	`

	db := u.trx.GetConn(ctx)
	result, err := db.ExecContext(ctx, query, params.ID)
	if err != nil {
		log.Error().Msgf("error deleting user: %v", err)
		return
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		log.Error().Msgf("user with id %s not found", params.ID)
		return constants.ErrorNotAffected
	}

	log.Debug().Msgf("deleted user with id: %s", params.ID)
	return
}

func (u *userRepository) Find(ctx context.Context, find model.UserFind) (users []model.User, err error) {
	query := `--sql
		SELECT * FROM users
	`

	db := u.trx.GetConn(ctx)
	cur, err := db.QueryxContext(ctx, query)
	if err != nil {
		log.Error().Msgf("")
		return
	}

	users = make([]model.User, 0)

	for cur.Next() {
		var each model.User
		if err := cur.StructScan(&each); err != nil {
			return nil, err
		}
	}

	return
}

func (u *userRepository) FindOne(ctx context.Context, find model.UserFind) (user model.User, err error) {
	query := `--sql
		SELECT * FROM users
	`

	db := u.trx.GetConn(ctx)

	err = db.QueryRowxContext(ctx, query).StructScan(&user)
	if err != nil {
		return
	}

	return
}

func (u *userRepository) Update(ctx context.Context, data model.UserUpdate) (err error) {
	query := `--sql
		UPDATE users SET email = $2, password = $3, status = $4 WHERE id = $1
	`
	id := data.Params.ID
	userData := data.Data

	db := u.trx.GetConn(ctx)
	result, err := db.ExecContext(
		ctx,
		query,
		id,
		userData.Email,
		userData.Password,
		userData.Status,
	)
	if err != nil {
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return
	}

	if affected == 0 {
		return constants.ErrorNotAffected
	}

	return
}
