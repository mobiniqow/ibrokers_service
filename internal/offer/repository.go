package offer

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

func (r *Repository) CreateOffer(item Offer) (Offer, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return Offer{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllOffers(limit, page int, filters []operators.FilterBlock) (items []Offer, count int64) {
    _query := helper.QueryBuilder(Offer{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateOffer(item Offer) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteOffer(item Offer) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindOfferById(id int) (Offer, error) {
    var item Offer
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return Offer{}, errors.New("not found offer")
    }
    return item, nil
}
