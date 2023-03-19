package rest

import (
	"github/Abraxas-365/bitacora/internal/auth"
	"github/Abraxas-365/bitacora/internal/leakerrs"
	"github/Abraxas-365/bitacora/pkg/report/application"
	"github/Abraxas-365/bitacora/pkg/report/domain/models"

	"github.com/gofiber/fiber/v2"
)

func ControllerFactory(fiberApp *fiber.App, app application.Application) {
	r := fiberApp.Group("/report")

	r.Post("/", auth.JWTProtected(), func(c *fiber.Ctx) error {
		report := models.Report{}
		userTokenData, err := auth.ExtractTokenMetadata(c)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		if err := c.BodyParser(&report); err != nil {
			err := leakerrs.GetError(err)
			c.Status(err.Code).JSON(err)
		}
		if err := app.Create(report.Constructor(userTokenData.Nickname)); err != nil {
			return c.Status(500).JSON(err)
		}

		return c.SendStatus(201)
	})

	r.Delete("/id=:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := app.Delete(id); err != nil {
			return c.Status(500).JSON(err)
		}

		return c.SendStatus(201)
	})
	r.Get("/", func(c *fiber.Ctx) error {
		criteria := models.SearchCriteria{}
		if err := c.BodyParser(&criteria); err != nil {
			err := leakerrs.GetError(err)
			c.Status(err.Code).JSON(err)
		}
		search, err := app.Get(criteria)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		return c.Status(200).JSON(search)
	})

}
