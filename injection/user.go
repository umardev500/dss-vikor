package injection

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/umardev500/spk/application/delivery"
	"github.com/umardev500/spk/application/repository"
	"github.com/umardev500/spk/application/usecase"
	"github.com/umardev500/spk/config"
)

func UserInject(router fiber.Router, db *sqlx.DB) {
	userRepo := repository.NewUserRepository(config.NewTransaction(db))
	uc := usecase.NewUserUsercase(userRepo)
	r := router.Group("user")
	delivery.NewUserDelivery(uc, r)
}
