package contract_type

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateContractType(item ContractType) (ContractType, error) {
    return s.Repository.CreateContractType(item)
}

func (s *Service) UpdateContractType(item ContractType) error {
    _, err := s.Repository.FindContractTypeById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateContractType(item)
}

func (s *Service) DeleteContractType(item ContractType) error {
    return s.Repository.DeleteContractType(item)
}

func (s *Service) GetAllContractTypes(limit, page int, filters []operators.FilterBlock) ([]ContractType, int64) {
    return s.Repository.GetAllContractTypes(limit, page, filters)
}
