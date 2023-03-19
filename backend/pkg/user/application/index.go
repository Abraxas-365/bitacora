package application

import (
	"github/Abraxas-365/bitacora/pkg/user/domain/models"
	"github/Abraxas-365/bitacora/pkg/user/domain/ports"
)

type Application interface {
	Create(new models.User) error
	LoginUser(username string, password string) (models.User, error)
	Delete(id interface{}) error
	Update(report models.User) error
	Get(key string, value interface{}) (models.Users, error)
}

type application struct {
	repo ports.UserRepository
}

func ApplicationFactory(repo ports.UserRepository) Application {
	return &application{
		repo,
	}
}
