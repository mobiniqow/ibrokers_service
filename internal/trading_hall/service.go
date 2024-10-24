package trading_hall

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateTradingHall(item TradingHall) (TradingHall, error) {
    return s.Repository.CreateTradingHall(item)
}

func (s *Service) UpdateTradingHall(item TradingHall) error {
    _, err := s.Repository.FindTradingHallById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateTradingHall(item)
}

func (s *Service) DeleteTradingHall(item TradingHall) error {
    return s.Repository.DeleteTradingHall(item)
}

func (s *Service) GetAllTradingHalls(limit, page int, filters []operators.FilterBlock) ([]TradingHall, int64) {
    return s.Repository.GetAllTradingHalls(limit, page, filters)
}
