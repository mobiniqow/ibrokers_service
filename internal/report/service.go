package report

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateReport(item Report) (Report, error) {
    return s.Repository.CreateReport(item)
}

func (s *Service) UpdateReport(item Report) error {
    _, err := s.Repository.FindReportById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateReport(item)
}

func (s *Service) DeleteReport(item Report) error {
    return s.Repository.DeleteReport(item)
}

func (s *Service) GetAllReports(limit, page int, filters []operators.FilterBlock) ([]Report, int64) {
    return s.Repository.GetAllReports(limit, page, filters)
}
