package application

import (
	"github/Abraxas-365/bitacora/internal/auth"
	"github/Abraxas-365/bitacora/pkg/auth/domain/models"
)

func (a *app) LoginUser(login models.Login) (string, error) {
	user, err := a.userApp.LoginUser(login.Username, login.Password)
	if err != nil {
		return "", err
	}

	jwt, err := auth.GereteToken(user.Nickname)
	if err != nil {
		return "", err
	}

	return jwt, nil
}
