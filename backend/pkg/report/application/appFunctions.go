package application

import (
	"github/Abraxas-365/bitacora/pkg/report/domain/models"
	"log"

	"github.com/google/uuid"
)

func (app *application) Create(new models.Report) error {
	if err := app.repo.Create(new); err != nil {
		return err
	}

	if err := app.elastic.Index(new); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (app *application) Delete(id interface{}) error {

	uuid, err := uuid.Parse(id.(string))
	if err != nil {
		return err
	}
	report, err := app.repo.Get("_id", uuid)
	if err != nil {
		return err
	}
	report[0].Delete = true
	if err := app.repo.Update(report[0]); err != nil {
		return err
	}
	_ = app.elastic.Delete(id.(string))

	return nil
}

func (app *application) Update(report models.Report) error {
	return app.repo.Update(report)
}

func (app *application) Get(criteria models.SearchCriteria) (models.Reports, error) {
	return app.elastic.Search(criteria)
}
