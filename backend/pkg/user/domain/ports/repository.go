package ports

import "github/Abraxas-365/bitacora/pkg/user/domain/models"

type UserRepository interface {
	Create(new models.User) error
	Delete(id interface{}) error
	Update(report models.User) error
	Get(key string, value interface{}) (models.Users, error)
	GetUser(key string, value interface{}) (models.User, bool, error)
}
