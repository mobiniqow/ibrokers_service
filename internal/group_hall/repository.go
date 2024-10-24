package group_hall

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

func (r *Repository) CreateGroupHall(item GroupHall) (GroupHall, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return GroupHall{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllGroupHalls(limit, page int, filters []operators.FilterBlock) (items []GroupHall, count int64) {
    _query := helper.QueryBuilder(GroupHall{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateGroupHall(item GroupHall) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteGroupHall(item GroupHall) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindGroupHallById(id int) (GroupHall, error) {
    var item GroupHall
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return GroupHall{}, errors.New("not found grouphall")
    }
    return item, nil
}
