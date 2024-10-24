package offer_mod

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

func (r *Repository) CreateOfferMod(item OfferMod) (OfferMod, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return OfferMod{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllOfferMods(limit, page int, filters []operators.FilterBlock) (items []OfferMod, count int64) {
    _query := helper.QueryBuilder(OfferMod{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateOfferMod(item OfferMod) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteOfferMod(item OfferMod) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindOfferModById(id int) (OfferMod, error) {
    var item OfferMod
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return OfferMod{}, errors.New("not found offermod")
    }
    return item, nil
}
