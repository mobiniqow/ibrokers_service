package buy_method

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

func (r *Repository) CreateBuyMethod(item BuyMethod) (BuyMethod, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return BuyMethod{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllBuyMethods(limit, page int, filters []operators.FilterBlock) (items []BuyMethod, count int64) {
    _query := helper.QueryBuilder(BuyMethod{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateBuyMethod(item BuyMethod) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteBuyMethod(item BuyMethod) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindBuyMethodById(id int) (BuyMethod, error) {
    var item BuyMethod
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return BuyMethod{}, errors.New("not found buymethod")
    }
    return item, nil
}
