package measure_unit

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateMeasureUnit(item MeasureUnit) (MeasureUnit, error) {
    return s.Repository.CreateMeasureUnit(item)
}

func (s *Service) UpdateMeasureUnit(item MeasureUnit) error {
    _, err := s.Repository.FindMeasureUnitById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateMeasureUnit(item)
}

func (s *Service) DeleteMeasureUnit(item MeasureUnit) error {
    return s.Repository.DeleteMeasureUnit(item)
}

func (s *Service) GetAllMeasureUnits(limit, page int, filters []operators.FilterBlock) ([]MeasureUnit, int64) {
    return s.Repository.GetAllMeasureUnits(limit, page, filters)
}
