package repository

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/umardev500/spk/config"
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

func (u *userRepository) Delete(context.Context, model.UserParams) (err error) {
	panic("unimplemented")
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
