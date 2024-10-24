package packaging_type

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

func (r *Repository) CreatePackagingType(item PackagingType) (PackagingType, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return PackagingType{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllPackagingTypes(limit, page int, filters []operators.FilterBlock) (items []PackagingType, count int64) {
    _query := helper.QueryBuilder(PackagingType{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdatePackagingType(item PackagingType) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeletePackagingType(item PackagingType) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindPackagingTypeById(id int) (PackagingType, error) {
    var item PackagingType
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return PackagingType{}, errors.New("not found packagingtype")
    }
    return item, nil
}
