package application

import (
	"errors"
	"fmt"
	"github/Abraxas-365/bitacora/internal/leakerrs"
	"github/Abraxas-365/bitacora/pkg/user/domain/models"
)

func (app *application) Create(new models.User) error {
	_, exist, err := app.repo.GetUser("nickname", new.Nickname)
	if err != nil {
		return err
	}
	if exist {
		return errors.New(leakerrs.DocumentExist)
	}

	return app.repo.Create(new)
}

func (app *application) Delete(id interface{}) error {
	return app.repo.Delete(id)
}

func (app *application) Update(report models.User) error {
	return app.repo.Update(report)
}

func (app *application) Get(key string, value interface{}) (models.Users, error) {
	return app.repo.Get(key, value)
}

func (a *application) LoginUser(username string, password string) (models.User, error) {
	//check if user exist
	user, exist, err := a.repo.GetUser("nikname", username)
	fmt.Println(user)
	if err != nil {
		return models.User{}, err
	}
	if !exist {
		return models.User{}, errors.New(leakerrs.DocumentNotFound)
	}

	if user.Password != password {
		return models.User{}, errors.New(leakerrs.DocumentNotFound)
	}

	return user, nil
}
