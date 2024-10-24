package hall_menu_sub_group

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateHallMenuSubGroup(item HallMenuSubGroup) (HallMenuSubGroup, error) {
    return s.Repository.CreateHallMenuSubGroup(item)
}

func (s *Service) UpdateHallMenuSubGroup(item HallMenuSubGroup) error {
    _, err := s.Repository.FindHallMenuSubGroupById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateHallMenuSubGroup(item)
}

func (s *Service) DeleteHallMenuSubGroup(item HallMenuSubGroup) error {
    return s.Repository.DeleteHallMenuSubGroup(item)
}

func (s *Service) GetAllHallMenuSubGroups(limit, page int, filters []operators.FilterBlock) ([]HallMenuSubGroup, int64) {
    return s.Repository.GetAllHallMenuSubGroups(limit, page, filters)
}
