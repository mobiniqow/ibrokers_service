package offer

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
	Repository Repository
}

func (s *Service) CreateOffer(item Offer) (Offer, error) {
	return s.Repository.CreateOffer(item)
}

func (s *Service) UpdateOffer(item Offer) error {
	_, err := s.Repository.FindOfferById(item.ID)
	if err != nil {
		return err
	}
	return s.Repository.UpdateOffer(item)
}

func (s *Service) DeleteOffer(item Offer) error {
	return s.Repository.DeleteOffer(item)
}

func (s *Service) GetAllOffers(limit, page int, filters []operators.FilterBlock) ([]Offer, int64) {
	return s.Repository.GetAllOffers(limit, page, filters)
}
