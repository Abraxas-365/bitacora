package rest

import (
	"fmt"
	"github/Abraxas-365/bitacora/internal/auth"
	"github/Abraxas-365/bitacora/internal/myerror"
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
			return c.Status(500).JSON(myerror.Wrap(err, 500).ToJson())
		}

		if err := c.BodyParser(&report); err != nil {
			c.Status(500).JSON(err)
		}
		if err := app.Create(report.Constructor(userTokenData.Nickname)); err != nil {
			return c.Status(500).JSON(err)
		}

		return c.SendStatus(201)
	})

	r.Delete("/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		fmt.Println("id", id)
		if err := app.Delete(id); err != nil {
			return c.Status(err.(*myerror.MyError).Code()).JSON(err.(*myerror.MyError).ToJson())
		}

		return c.SendStatus(200)
	})
	r.Get("/", func(c *fiber.Ctx) error {
		criteria := models.SearchCriteria{}
		if err := c.BodyParser(&criteria); err != nil {
			return c.Status(500).JSON(myerror.Wrap(err, 500).ToJson())
		}
		search, err := app.Get(criteria)
		if err != nil {
			return c.Status(err.(*myerror.MyError).Code()).JSON(err.(*myerror.MyError).ToJson())
		}
		return c.Status(200).JSON(search)
	})

	r.Put("/:id", auth.JWTProtected(), func(c *fiber.Ctx) error {
		report := models.Report{}
		id := c.Params("id")
		_, err := auth.ExtractTokenMetadata(c)
		if err != nil {
			return c.Status(500).JSON(myerror.Wrap(err, 500).ToJson())
		}

		if err := c.BodyParser(&report); err != nil {
			return c.Status(500).JSON(myerror.Wrap(err, 500).ToJson())
		}
		if err := app.Update(report, id); err != nil {
			return c.Status(err.(*myerror.MyError).Code()).JSON(err.(*myerror.MyError).ToJson())
		}
		return c.SendStatus(200)
	})

}
