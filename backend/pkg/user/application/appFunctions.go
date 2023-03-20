package application

import (
	"errors"
	"fmt"
	"github/Abraxas-365/bitacora/internal/myerror"
	"github/Abraxas-365/bitacora/pkg/user/domain/models"
)

func (app *application) Create(new models.User) error {
	_, exist, err := app.repo.GetUser("nickname", new.Nickname)
	if err != nil {
		return err
	}
	if exist {
		return myerror.Wrap(errors.New("user already exist"), 400)
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
	user, _, err := a.repo.GetUser("nikname", username)
	fmt.Println(user)
	if err != nil {
		return models.User{}, err
	}

	if user.Password != password {
		return models.User{}, myerror.Wrap(errors.New("invalid password"), 401)
	}

	return user, nil
}
