package manufacturers

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateManufacturers(item Manufacturers) (Manufacturers, error) {
    return s.Repository.CreateManufacturers(item)
}

func (s *Service) UpdateManufacturers(item Manufacturers) error {
    _, err := s.Repository.FindManufacturersById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateManufacturers(item)
}

func (s *Service) DeleteManufacturers(item Manufacturers) error {
    return s.Repository.DeleteManufacturers(item)
}

func (s *Service) GetAllManufacturerss(limit, page int, filters []operators.FilterBlock) ([]Manufacturers, int64) {
    return s.Repository.GetAllManufacturerss(limit, page, filters)
}
