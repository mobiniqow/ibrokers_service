package supplier

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

func (r *Repository) CreateSupplier(item Supplier) (Supplier, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return Supplier{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllSuppliers(limit, page int, filters []operators.FilterBlock) (items []Supplier, count int64) {
    _query := helper.QueryBuilder(Supplier{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateSupplier(item Supplier) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteSupplier(item Supplier) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindSupplierById(id int) (Supplier, error) {
    var item Supplier
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return Supplier{}, errors.New("not found supplier")
    }
    return item, nil
}
