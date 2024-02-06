package repository

import (
	"context"
	"fmt"

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

func (u *userRepository) Create(ctx context.Context, data model.UserCreate) (err error) {
	query := `--sql
		INSERT INTO users (id, email, password, status)
		VALUES ($1, $2, $3, $4);
	`

	db := u.trx.GetConn(ctx)
	user := data.Data

	_, err = db.ExecContext(ctx, query, user.ID, user.Email, user.Password, user.Status)
	if err != nil {
		return
	}

	return
}

func (u *userRepository) Delete(ctx context.Context, params model.UserParams) (err error) {
	query := `--sql
		DELETE FROM users WHERE id = $1;
	`

	db := u.trx.GetConn(ctx)
	result, err := db.ExecContext(ctx, query, params.ID)
	if err != nil {
		return
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return constants.ErrorNotAffected
	}

	return
}

func (u *userRepository) Find(ctx context.Context, find model.UserFind) (users []model.User, err error) {
	query := `--sql
		SELECT * FROM users
	`

	db := u.trx.GetConn(ctx)
	cur, err := db.QueryxContext(ctx, query)
	if err != nil {
		return
	}

	users = make([]model.User, 0)

	for cur.Next() {
		var each model.User
		if err := cur.StructScan(&each); err != nil {
			return nil, err
		}

		users = append(users, each)
	}

	return
}

func (u *userRepository) FindOne(ctx context.Context, find model.UserFind) (user model.User, err error) {
	query := `--sql
		SELECT * FROM users WHERE id = $1
	`

	db := u.trx.GetConn(ctx)
	id := find.ID
	fmt.Println("id:", id)

	err = db.QueryRowxContext(ctx, query, id).StructScan(&user)
	if err != nil {
		fmt.Println(err)
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
