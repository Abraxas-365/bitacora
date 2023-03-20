package application

import (
	"fmt"
	"github/Abraxas-365/bitacora/pkg/report/domain/models"
	"log"

	"github.com/google/uuid"
)

func (app *application) Create(new models.Report) error {
	if err := app.repo.Create(new); err != nil {
		return err
	}

	if err := app.elastic.Index(new); err != nil {
		log.Println(err)
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
	fmt.Println(report)
	report[0].Delete = true
	if err := app.repo.Update(report[0]); err != nil {
		return err
	}
	_ = app.elastic.Delete(id.(string))

	return nil
}

func (app *application) Update(ToUpdate models.Report, id string) error {

	_ = app.elastic.Delete(id)

	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	reports, err := app.repo.Get("_id", uuid)
	if err != nil {
		return err
	}
	//Check which fields are updated
	if ToUpdate.Title != "" {
		reports[0].Title = ToUpdate.Title
	}
	if ToUpdate.Data != "" {
		reports[0].Data = ToUpdate.Data
	}
	if ToUpdate.Tags != nil {
		reports[0].Tags = ToUpdate.Tags
	}
	if ToUpdate.Error != "" {
		reports[0].Error = ToUpdate.Error
	}
	if ToUpdate.Solution != "" {
		reports[0].Solution = ToUpdate.Solution
	}

	if err := app.repo.Update(reports[0]); err != nil {
		return err
	}

	if err := app.elastic.Index(reports[0]); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (app *application) Get(criteria models.SearchCriteria) (models.Reports, error) {
	return app.elastic.Search(criteria)
}
