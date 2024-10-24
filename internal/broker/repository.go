package broker

import (
    "errors"
    "ibrokers_service/pkg/helper"
    "ibrokers_service/pkg/middleware/filter/operators"
    "ibrokers_service/pkg/middleware/pagination"

    "gorm.io/gorm"
)

type Repository struct {
    DB *gorm.DB
}

func (r *Repository) CreateBroker(item Broker) (Broker, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return Broker{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllBrokers(limit, page int, filters []operators.FilterBlock) (items []Broker, count int64) {
    _query := helper.QueryBuilder(Broker{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateBroker(item Broker) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteBroker(item Broker) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindBrokerById(id int) (Broker, error) {
    var item Broker
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return Broker{}, errors.New("not found broker")
    }
    return item, nil
}
