package rest

import (
	"fmt"
	"github/Abraxas-365/bitacora/internal/myerror"
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
			return c.Status(500).JSON(myerror.Wrap(err, 500).ToJson())

		}

		token, err := app.LoginUser(login)
		if err != nil {
			return c.Status(500).JSON(myerror.Wrap(err, 500).ToJson())
		}

		return c.Status(200).JSON(token)
	})
}
