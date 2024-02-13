package injection

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/umardev500/spk/application/delivery"
	"github.com/umardev500/spk/application/repository"
	"github.com/umardev500/spk/application/usecase"
	"github.com/umardev500/spk/config"
)

func RoleInject(router fiber.Router, db *sqlx.DB) {
	roleRepo := repository.NewRoleRepository(config.NewTransaction(db))
	uc := usecase.NewRoleUsecase(roleRepo)
	r := router.Group("roles")
	delivery.NewRoleDelivery(uc, r)
}
