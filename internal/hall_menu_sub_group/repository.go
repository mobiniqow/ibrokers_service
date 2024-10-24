package hall_menu_sub_group

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

func (r *Repository) CreateHallMenuSubGroup(item HallMenuSubGroup) (HallMenuSubGroup, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return HallMenuSubGroup{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllHallMenuSubGroups(limit, page int, filters []operators.FilterBlock) (items []HallMenuSubGroup, count int64) {
    _query := helper.QueryBuilder(HallMenuSubGroup{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateHallMenuSubGroup(item HallMenuSubGroup) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteHallMenuSubGroup(item HallMenuSubGroup) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindHallMenuSubGroupById(id int) (HallMenuSubGroup, error) {
    var item HallMenuSubGroup
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return HallMenuSubGroup{}, errors.New("not found hallmenusubgroup")
    }
    return item, nil
}
