package offer_mod

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateOfferMod(item OfferMod) (OfferMod, error) {
    return s.Repository.CreateOfferMod(item)
}

func (s *Service) UpdateOfferMod(item OfferMod) error {
    _, err := s.Repository.FindOfferModById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateOfferMod(item)
}

func (s *Service) DeleteOfferMod(item OfferMod) error {
    return s.Repository.DeleteOfferMod(item)
}

func (s *Service) GetAllOfferMods(limit, page int, filters []operators.FilterBlock) ([]OfferMod, int64) {
    return s.Repository.GetAllOfferMods(limit, page, filters)
}
