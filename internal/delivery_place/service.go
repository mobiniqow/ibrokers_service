package delivery_place

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateDeliveryPlace(item DeliveryPlace) (DeliveryPlace, error) {
    return s.Repository.CreateDeliveryPlace(item)
}

func (s *Service) UpdateDeliveryPlace(item DeliveryPlace) error {
    _, err := s.Repository.FindDeliveryPlaceById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateDeliveryPlace(item)
}

func (s *Service) DeleteDeliveryPlace(item DeliveryPlace) error {
    return s.Repository.DeleteDeliveryPlace(item)
}

func (s *Service) GetAllDeliveryPlaces(limit, page int, filters []operators.FilterBlock) ([]DeliveryPlace, int64) {
    return s.Repository.GetAllDeliveryPlaces(limit, page, filters)
}
