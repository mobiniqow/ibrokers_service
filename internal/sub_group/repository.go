package sub_group

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

func (r *Repository) CreateSubGroup(item SubGroup) (SubGroup, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return SubGroup{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllSubGroups(limit, page int, filters []operators.FilterBlock) (items []SubGroup, count int64) {
    _query := helper.QueryBuilder(SubGroup{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateSubGroup(item SubGroup) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteSubGroup(item SubGroup) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindSubGroupById(id int) (SubGroup, error) {
    var item SubGroup
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return SubGroup{}, errors.New("not found subgroup")
    }
    return item, nil
}
