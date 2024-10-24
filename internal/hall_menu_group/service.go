package hall_menu_group

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateHallMenuGroup(item HallMenuGroup) (HallMenuGroup, error) {
    return s.Repository.CreateHallMenuGroup(item)
}

func (s *Service) UpdateHallMenuGroup(item HallMenuGroup) error {
    _, err := s.Repository.FindHallMenuGroupById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateHallMenuGroup(item)
}

func (s *Service) DeleteHallMenuGroup(item HallMenuGroup) error {
    return s.Repository.DeleteHallMenuGroup(item)
}

func (s *Service) GetAllHallMenuGroups(limit, page int, filters []operators.FilterBlock) ([]HallMenuGroup, int64) {
    return s.Repository.GetAllHallMenuGroups(limit, page, filters)
}
