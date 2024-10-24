package packaging_type

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreatePackagingType(item PackagingType) (PackagingType, error) {
    return s.Repository.CreatePackagingType(item)
}

func (s *Service) UpdatePackagingType(item PackagingType) error {
    _, err := s.Repository.FindPackagingTypeById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdatePackagingType(item)
}

func (s *Service) DeletePackagingType(item PackagingType) error {
    return s.Repository.DeletePackagingType(item)
}

func (s *Service) GetAllPackagingTypes(limit, page int, filters []operators.FilterBlock) ([]PackagingType, int64) {
    return s.Repository.GetAllPackagingTypes(limit, page, filters)
}
