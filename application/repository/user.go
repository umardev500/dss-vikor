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
		VALUES (?, ?, ?, ?);
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
		DELETE FROM users WHERE id = ?;
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

func (u *userRepository) Find(context.Context, model.UserFind) (users []model.User, err error) {
	panic("unimplemented")
}

func (u *userRepository) FindOne(context.Context, model.UserFind) (user model.User, err error) {
	panic("unimplemented")
}

func (u *userRepository) Update(context.Context, model.UserUpdate) (err error) {
	panic("unimplemented")
}
