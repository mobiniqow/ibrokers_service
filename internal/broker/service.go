package broker

import "ibrokers_service/pkg/middleware/filter/operators"

type Service struct {
    Repository Repository
}

func (s *Service) CreateBroker(item Broker) (Broker, error) {
    return s.Repository.CreateBroker(item)
}

func (s *Service) UpdateBroker(item Broker) error {
    _, err := s.Repository.FindBrokerById(item.Id)
    if err != nil {
        return err
    }
    return s.Repository.UpdateBroker(item)
}

func (s *Service) DeleteBroker(item Broker) error {
    return s.Repository.DeleteBroker(item)
}

func (s *Service) GetAllBrokers(limit, page int, filters []operators.FilterBlock) ([]Broker, int64) {
    return s.Repository.GetAllBrokers(limit, page, filters)
}
