package rest

import (
	"github/Abraxas-365/bitacora/internal/myerror"
	"github/Abraxas-365/bitacora/pkg/user/application"
	"github/Abraxas-365/bitacora/pkg/user/domain/models"

	"github.com/gofiber/fiber/v2"
)

func ControllerFactory(fiberApp *fiber.App, app application.Application) {
	r := fiberApp.Group("/user")

	r.Post("/", func(c *fiber.Ctx) error {
		user := models.User{}

		if err := c.BodyParser(&user); err != nil {
			return c.Status(500).JSON(myerror.Wrap(err, 500).ToJson())
		}
		if err := app.Create(user.Constructor()); err != nil {
			return c.SendStatus(500)
		}

		return c.SendStatus(201)
	})

	r.Delete("/id=:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := app.Delete(id); err != nil {
			return c.Status(500).JSON(myerror.Wrap(err, 500).ToJson())
		}

		return c.SendStatus(200)
	})

}
