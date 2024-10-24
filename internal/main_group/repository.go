package main_group

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

func (r *Repository) CreateMainGroup(item MainGroup) (MainGroup, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return MainGroup{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllMainGroups(limit, page int, filters []operators.FilterBlock) (items []MainGroup, count int64) {
    _query := helper.QueryBuilder(MainGroup{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateMainGroup(item MainGroup) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteMainGroup(item MainGroup) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindMainGroupById(id int) (MainGroup, error) {
    var item MainGroup
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return MainGroup{}, errors.New("not found maingroup")
    }
    return item, nil
}
