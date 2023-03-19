package report

import (
	"github/Abraxas-365/bitacora/internal/mongo"
	"github/Abraxas-365/bitacora/pkg/report/application"
	"github/Abraxas-365/bitacora/pkg/report/infrastructure/elasticrepo"
	"github/Abraxas-365/bitacora/pkg/report/infrastructure/repository"
	"github/Abraxas-365/bitacora/pkg/report/infrastructure/rest"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gofiber/fiber/v2"
)

func ModuleFactory(appFiber *fiber.App, mongo mongo.MongoConnection, elastic *elasticsearch.Client) {
	repo := repository.RepositoryFactory(mongo, "reports")
	elasticRepo := elasticrepo.ElasticRepositoryFactory(elastic, "reports")
	application := application.ApplicationFactory(repo, elasticRepo)
	rest.ControllerFactory(appFiber, application)
}
