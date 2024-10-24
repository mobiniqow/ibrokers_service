package report

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

func (r *Repository) CreateReport(item Report) (Report, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return Report{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllReports(limit, page int, filters []operators.FilterBlock) (items []Report, count int64) {
    _query := helper.QueryBuilder(Report{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateReport(item Report) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteReport(item Report) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindReportById(id int) (Report, error) {
    var item Report
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return Report{}, errors.New("not found report")
    }
    return item, nil
}
