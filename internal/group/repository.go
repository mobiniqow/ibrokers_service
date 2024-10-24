package group

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

func (r *Repository) CreateGroup(item Group) (Group, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return Group{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllGroups(limit, page int, filters []operators.FilterBlock) (items []Group, count int64) {
    _query := helper.QueryBuilder(Group{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateGroup(item Group) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteGroup(item Group) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindGroupById(id int) (Group, error) {
    var item Group
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return Group{}, errors.New("not found group")
    }
    return item, nil
}
