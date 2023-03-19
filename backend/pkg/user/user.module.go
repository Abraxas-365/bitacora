package user

import (
	"github/Abraxas-365/bitacora/internal/mongo"
	"github/Abraxas-365/bitacora/pkg/user/application"
	"github/Abraxas-365/bitacora/pkg/user/infastrucutre/repository"
	"github/Abraxas-365/bitacora/pkg/user/infastrucutre/rest"

	"github.com/gofiber/fiber/v2"
)

func ModuleFactory(appFiber *fiber.App, mongo mongo.MongoConnection) application.Application {
	repo := repository.RepositoryFactory(mongo, "users")
	application := application.ApplicationFactory(repo)
	rest.ControllerFactory(appFiber, application)

	return application
}
