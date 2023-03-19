package repository

import (
	"time"

	mclient "github/Abraxas-365/bitacora/internal/mongo"
	"github/Abraxas-365/bitacora/pkg/user/domain/ports"

	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	client     *mongo.Client
	database   string
	collection string
	timeout    time.Duration
}

func RepositoryFactory(mongoConnection mclient.MongoConnection, collection string) ports.UserRepository {
	return &repo{
		client:     mongoConnection.Client,
		collection: collection,
		database:   mongoConnection.Database,
		timeout:    mongoConnection.Timeout,
	}

}
