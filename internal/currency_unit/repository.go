package currency_unit

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

func (r *Repository) CreateCurrencyUnit(item CurrencyUnit) (CurrencyUnit, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return CurrencyUnit{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllCurrencyUnits(limit, page int, filters []operators.FilterBlock) (items []CurrencyUnit, count int64) {
    _query := helper.QueryBuilder(CurrencyUnit{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateCurrencyUnit(item CurrencyUnit) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteCurrencyUnit(item CurrencyUnit) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindCurrencyUnitById(id int) (CurrencyUnit, error) {
    var item CurrencyUnit
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return CurrencyUnit{}, errors.New("not found currencyunit")
    }
    return item, nil
}
