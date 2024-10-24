package trading_hall

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

func (r *Repository) CreateTradingHall(item TradingHall) (TradingHall, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return TradingHall{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllTradingHalls(limit, page int, filters []operators.FilterBlock) (items []TradingHall, count int64) {
    _query := helper.QueryBuilder(TradingHall{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateTradingHall(item TradingHall) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteTradingHall(item TradingHall) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindTradingHallById(id int) (TradingHall, error) {
    var item TradingHall
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return TradingHall{}, errors.New("not found tradinghall")
    }
    return item, nil
}
