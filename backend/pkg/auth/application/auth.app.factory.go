package application

import (
	"github/Abraxas-365/bitacora/pkg/auth/domain/models"
	user "github/Abraxas-365/bitacora/pkg/user/application"
)

type Application interface {
	LoginUser(login models.Login) (string, error)
}

type app struct {
	userApp user.Application
}

func AppFactory(userApp user.Application) Application {
	return &app{
		userApp: userApp,
	}
}
