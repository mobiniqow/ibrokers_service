package hall_menu_group

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

func (r *Repository) CreateHallMenuGroup(item HallMenuGroup) (HallMenuGroup, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return HallMenuGroup{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllHallMenuGroups(limit, page int, filters []operators.FilterBlock) (items []HallMenuGroup, count int64) {
    _query := helper.QueryBuilder(HallMenuGroup{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateHallMenuGroup(item HallMenuGroup) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteHallMenuGroup(item HallMenuGroup) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindHallMenuGroupById(id int) (HallMenuGroup, error) {
    var item HallMenuGroup
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return HallMenuGroup{}, errors.New("not found hallmenugroup")
    }
    return item, nil
}
