package application

import (
	"github/Abraxas-365/bitacora/pkg/report/domain/models"
	"github/Abraxas-365/bitacora/pkg/report/domain/ports"
)

type Application interface {
	Create(new models.Report) error
	Delete(id interface{}) error
	Update(report models.Report) error
	Get(models.SearchCriteria) (models.Reports, error)
}

type application struct {
	repo    ports.ReportRepository
	elastic ports.ReportSearchRepository
}

func ApplicationFactory(repo ports.ReportRepository, elastic ports.ReportSearchRepository) Application {
	return &application{
		repo,
		elastic,
	}
}
