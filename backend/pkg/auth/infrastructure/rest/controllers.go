package rest

import (
	"fmt"
	"github/Abraxas-365/bitacora/internal/leakerrs"
	"github/Abraxas-365/bitacora/pkg/auth/application"
	"github/Abraxas-365/bitacora/pkg/auth/domain/models"

	"github.com/gofiber/fiber/v2"
)

func Routes(appFiber *fiber.App, app application.Application) {
	auth := appFiber.Group("/auth")

	auth.Post("/login", func(c *fiber.Ctx) error {
		fmt.Println("LOGINNN")
		login := models.Login{}
		if err := c.BodyParser(&login); err != nil {
			err := leakerrs.GetError(err)
			c.Status(err.Code).JSON(err)
		}

		token, err := app.LoginUser(login)
		if err != nil {
			err := leakerrs.GetError(err)
			return c.Status(err.Code).JSON(err)
		}

		return c.Status(201).JSON(token)
	})
}
