package buy_method

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateBuyMethod(item BuyMethod) (BuyMethod, error) {
    return s.Repository.CreateBuyMethod(item)
}

func (s *Service) UpdateBuyMethod(item BuyMethod) error {
    _, err := s.Repository.FindBuyMethodById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateBuyMethod(item)
}

func (s *Service) DeleteBuyMethod(item BuyMethod) error {
    return s.Repository.DeleteBuyMethod(item)
}

func (s *Service) GetAllBuyMethods(limit, page int, filters []operators.FilterBlock) ([]BuyMethod, int64) {
    return s.Repository.GetAllBuyMethods(limit, page, filters)
}
