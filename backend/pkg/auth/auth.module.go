package auth

import (
	"github/Abraxas-365/bitacora/pkg/auth/application"
	"github/Abraxas-365/bitacora/pkg/auth/infrastructure/rest"
	user "github/Abraxas-365/bitacora/pkg/user/application"

	"github.com/gofiber/fiber/v2"
)

func ModuleFactory(app *fiber.App, Userapp user.Application) {
	application := application.AppFactory(Userapp)
	rest.Routes(app, application)

}
