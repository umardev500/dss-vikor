package injection

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/umardev500/spk/application/delivery"
	"github.com/umardev500/spk/application/repository"
	"github.com/umardev500/spk/application/usecase"
	"github.com/umardev500/spk/config"
)

func CriteriaInject(router fiber.Router, db *sqlx.DB, v *validator.Validate) {
	criteriaRepo := repository.NewCriteriaRepository(config.NewTransaction(db))
	uc := usecase.NewCriteriaUsecase(criteriaRepo)
	r := router.Group("criterias")
	delivery.NewCriteriaDelivery(uc, r, v)
}
