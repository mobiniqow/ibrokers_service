package main_group

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateMainGroup(item MainGroup) (MainGroup, error) {
    return s.Repository.CreateMainGroup(item)
}

func (s *Service) UpdateMainGroup(item MainGroup) error {
    _, err := s.Repository.FindMainGroupById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateMainGroup(item)
}

func (s *Service) DeleteMainGroup(item MainGroup) error {
    return s.Repository.DeleteMainGroup(item)
}

func (s *Service) GetAllMainGroups(limit, page int, filters []operators.FilterBlock) ([]MainGroup, int64) {
    return s.Repository.GetAllMainGroups(limit, page, filters)
}
