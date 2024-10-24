package currency_unit

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateCurrencyUnit(item CurrencyUnit) (CurrencyUnit, error) {
    return s.Repository.CreateCurrencyUnit(item)
}

func (s *Service) UpdateCurrencyUnit(item CurrencyUnit) error {
    _, err := s.Repository.FindCurrencyUnitById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateCurrencyUnit(item)
}

func (s *Service) DeleteCurrencyUnit(item CurrencyUnit) error {
    return s.Repository.DeleteCurrencyUnit(item)
}

func (s *Service) GetAllCurrencyUnits(limit, page int, filters []operators.FilterBlock) ([]CurrencyUnit, int64) {
    return s.Repository.GetAllCurrencyUnits(limit, page, filters)
}
