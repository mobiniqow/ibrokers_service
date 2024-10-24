package sub_group

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateSubGroup(item SubGroup) (SubGroup, error) {
    return s.Repository.CreateSubGroup(item)
}

func (s *Service) UpdateSubGroup(item SubGroup) error {
    _, err := s.Repository.FindSubGroupById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateSubGroup(item)
}

func (s *Service) DeleteSubGroup(item SubGroup) error {
    return s.Repository.DeleteSubGroup(item)
}

func (s *Service) GetAllSubGroups(limit, page int, filters []operators.FilterBlock) ([]SubGroup, int64) {
    return s.Repository.GetAllSubGroups(limit, page, filters)
}
