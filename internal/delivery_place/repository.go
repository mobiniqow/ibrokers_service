package delivery_place

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

func (r *Repository) CreateDeliveryPlace(item DeliveryPlace) (DeliveryPlace, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return DeliveryPlace{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllDeliveryPlaces(limit, page int, filters []operators.FilterBlock) (items []DeliveryPlace, count int64) {
    _query := helper.QueryBuilder(DeliveryPlace{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateDeliveryPlace(item DeliveryPlace) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteDeliveryPlace(item DeliveryPlace) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindDeliveryPlaceById(id int) (DeliveryPlace, error) {
    var item DeliveryPlace
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return DeliveryPlace{}, errors.New("not found deliveryplace")
    }
    return item, nil
}
