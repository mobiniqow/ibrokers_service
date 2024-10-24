package offer_type

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateOfferType(item OfferType) (OfferType, error) {
    return s.Repository.CreateOfferType(item)
}

func (s *Service) UpdateOfferType(item OfferType) error {
    _, err := s.Repository.FindOfferTypeById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateOfferType(item)
}

func (s *Service) DeleteOfferType(item OfferType) error {
    return s.Repository.DeleteOfferType(item)
}

func (s *Service) GetAllOfferTypes(limit, page int, filters []operators.FilterBlock) ([]OfferType, int64) {
    return s.Repository.GetAllOfferTypes(limit, page, filters)
}
