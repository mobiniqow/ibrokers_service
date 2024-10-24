package group_hall

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateGroupHall(item GroupHall) (GroupHall, error) {
    return s.Repository.CreateGroupHall(item)
}

func (s *Service) UpdateGroupHall(item GroupHall) error {
    _, err := s.Repository.FindGroupHallById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateGroupHall(item)
}

func (s *Service) DeleteGroupHall(item GroupHall) error {
    return s.Repository.DeleteGroupHall(item)
}

func (s *Service) GetAllGroupHalls(limit, page int, filters []operators.FilterBlock) ([]GroupHall, int64) {
    return s.Repository.GetAllGroupHalls(limit, page, filters)
}
