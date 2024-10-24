package settlement

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

func (r *Repository) CreateSettlement(item Settlement) (Settlement, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return Settlement{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllSettlements(limit, page int, filters []operators.FilterBlock) (items []Settlement, count int64) {
    _query := helper.QueryBuilder(Settlement{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateSettlement(item Settlement) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteSettlement(item Settlement) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindSettlementById(id int) (Settlement, error) {
    var item Settlement
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return Settlement{}, errors.New("not found settlement")
    }
    return item, nil
}
