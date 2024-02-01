package migrations

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func Migrate(db *sqlx.DB) {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatal().Msgf("failed to create postgres driver: %v", err)
		return
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations/schemas",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatal().Msgf("failed to create migrate instance: %v", err)
		return
	}

	log.Info().Msg("migrating...")
	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Warn().Msg(err.Error())
		} else {

			log.Fatal().Msgf("failed to migrate: %v", err)
		}
	}
	log.Info().Msg("migrations completed")
}
