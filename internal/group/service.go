package group

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateGroup(item Group) (Group, error) {
    return s.Repository.CreateGroup(item)
}

func (s *Service) UpdateGroup(item Group) error {
    _, err := s.Repository.FindGroupById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateGroup(item)
}

func (s *Service) DeleteGroup(item Group) error {
    return s.Repository.DeleteGroup(item)
}

func (s *Service) GetAllGroups(limit, page int, filters []operators.FilterBlock) ([]Group, int64) {
    return s.Repository.GetAllGroups(limit, page, filters)
}
