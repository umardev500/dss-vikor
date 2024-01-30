package config

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/spk/constants"
)

func NewPostgres(ctx context.Context) (db *sqlx.DB) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPass)

	start := time.Now()
	db, err := sqlx.ConnectContext(ctx, "postgres", dsn)
	if err != nil {
		log.Fatal().Msgf("error connecting to postgres: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal().Msgf("error pinging postgres: %v", err)
	}

	duration := time.Since(start)
	msg := fmt.Sprintf("Connected to Postgres \033[32m🎉 (\U000023F3 %s)\033[0m", duration)
	log.Info().Msg(msg)

	return
}

type Queries interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type DB interface {
	BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
	Queries
}

type Trx struct {
	db DB
}

// TxFn is a function that can be executed inside a transaction.
type TxFn func(context.Context) error

// NewTransaction creates a new transaction.
//
// It takes a PostgresDB as a parameter and returns a pointer to a Trx.
func NewTransaction(db DB) *Trx {
	return &Trx{
		db: db,
	}
}

// GetConn returns the database connection object.
//
// It takes a context.Context as a parameter and returns a Queries object.
func (t *Trx) GetConn(ctx context.Context) (db Queries) {
	log.Debug().Msgf("Get database connection")
	db, ok := ctx.Value(constants.CtxKeyTx).(Queries)
	if !ok {
		// log.Debug().Msgf("No database found in context, default is db used")
		return t.db
	}

	log.Debug().Msgf("Database found in context")

	return
}

// WithTransaction executes a function inside a transaction.
//
// ctx: The context in which the transaction should be executed.
// fn: The function to be executed inside the transaction.
// err: The error returned by the function, if any.
func (t *Trx) WithTransaction(ctx context.Context, fn TxFn) (err error) {
	log.Debug().Msgf("Begin transaction")
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			log.Error().Msgf("Rollback transaction: %s", err)
			tx.Rollback()
		} else {
			log.Debug().Msgf("Commit transaction")
			err = tx.Commit()
		}

		log.Debug().Msgf("End transaction")
	}()

	ctx = context.WithValue(ctx, constants.CtxKeyTx, tx)
	err = fn(ctx)

	return
}
