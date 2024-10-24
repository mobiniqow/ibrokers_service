package commodity

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

func (r *Repository) CreateCommodity(item Commodity) (Commodity, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return Commodity{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllCommoditys(limit, page int, filters []operators.FilterBlock) (items []Commodity, count int64) {
    _query := helper.QueryBuilder(Commodity{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateCommodity(item Commodity) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteCommodity(item Commodity) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindCommodityById(id int) (Commodity, error) {
    var item Commodity
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return Commodity{}, errors.New("not found commodity")
    }
    return item, nil
}
