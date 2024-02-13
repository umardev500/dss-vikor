package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/umardev500/spk/injection"
)

type Router struct {
	app *fiber.App
	db  *sqlx.DB
}

func NewRouter(app *fiber.App, db *sqlx.DB) Router {
	return Router{
		app: app,
		db:  db,
	}
}

func (r Router) Register() {
	api := r.app.Group("api")
	injection.UserInject(api, r.db)
	injection.RoleInject(api, r.db)
}
