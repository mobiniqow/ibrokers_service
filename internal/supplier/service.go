package supplier

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateSupplier(item Supplier) (Supplier, error) {
    return s.Repository.CreateSupplier(item)
}

func (s *Service) UpdateSupplier(item Supplier) error {
    _, err := s.Repository.FindSupplierById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateSupplier(item)
}

func (s *Service) DeleteSupplier(item Supplier) error {
    return s.Repository.DeleteSupplier(item)
}

func (s *Service) GetAllSuppliers(limit, page int, filters []operators.FilterBlock) ([]Supplier, int64) {
    return s.Repository.GetAllSuppliers(limit, page, filters)
}
