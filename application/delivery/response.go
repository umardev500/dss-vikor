package delivery

import (
	"github.com/gofiber/fiber/v2"
	"github.com/umardev500/spk/domain/model"
)

func OkResponse(c *fiber.Ctx, res model.Response) error {

	return c.JSON(res)
}
