package ports

import "github/Abraxas-365/bitacora/pkg/report/domain/models"

type ReportRepository interface {
	Create(new models.Report) error
	Delete(id interface{}) error
	Update(report models.Report) error
	Get(key string, value interface{}) (models.Reports, error)
}
