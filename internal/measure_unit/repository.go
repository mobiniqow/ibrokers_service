package measure_unit

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

func (r *Repository) CreateMeasureUnit(item MeasureUnit) (MeasureUnit, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return MeasureUnit{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllMeasureUnits(limit, page int, filters []operators.FilterBlock) (items []MeasureUnit, count int64) {
    _query := helper.QueryBuilder(MeasureUnit{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateMeasureUnit(item MeasureUnit) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteMeasureUnit(item MeasureUnit) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindMeasureUnitById(id int) (MeasureUnit, error) {
    var item MeasureUnit
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return MeasureUnit{}, errors.New("not found measureunit")
    }
    return item, nil
}
