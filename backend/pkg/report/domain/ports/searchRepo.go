package ports

import "github/Abraxas-365/bitacora/pkg/report/domain/models"

type ReportSearchRepository interface {
	Delete(id string) error
	Index(report models.Report) error
	Search(criteria models.SearchCriteria) (models.Reports, error)
}
