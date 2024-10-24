package settlement

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateSettlement(item Settlement) (Settlement, error) {
    return s.Repository.CreateSettlement(item)
}

func (s *Service) UpdateSettlement(item Settlement) error {
    _, err := s.Repository.FindSettlementById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateSettlement(item)
}

func (s *Service) DeleteSettlement(item Settlement) error {
    return s.Repository.DeleteSettlement(item)
}

func (s *Service) GetAllSettlements(limit, page int, filters []operators.FilterBlock) ([]Settlement, int64) {
    return s.Repository.GetAllSettlements(limit, page, filters)
}
