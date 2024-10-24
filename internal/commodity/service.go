package commodity

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateCommodity(item Commodity) (Commodity, error) {
    return s.Repository.CreateCommodity(item)
}

func (s *Service) UpdateCommodity(item Commodity) error {
    _, err := s.Repository.FindCommodityById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateCommodity(item)
}

func (s *Service) DeleteCommodity(item Commodity) error {
    return s.Repository.DeleteCommodity(item)
}

func (s *Service) GetAllCommoditys(limit, page int, filters []operators.FilterBlock) ([]Commodity, int64) {
    return s.Repository.GetAllCommoditys(limit, page, filters)
}
