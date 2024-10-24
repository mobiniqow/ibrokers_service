package manufacturers

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

func (r *Repository) CreateManufacturers(item Manufacturers) (Manufacturers, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return Manufacturers{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllManufacturerss(limit, page int, filters []operators.FilterBlock) (items []Manufacturers, count int64) {
    _query := helper.QueryBuilder(Manufacturers{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateManufacturers(item Manufacturers) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteManufacturers(item Manufacturers) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindManufacturersById(id int) (Manufacturers, error) {
    var item Manufacturers
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return Manufacturers{}, errors.New("not found manufacturers")
    }
    return item, nil
}
