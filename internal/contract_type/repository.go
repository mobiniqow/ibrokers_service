package contract_type

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

func (r *Repository) CreateContractType(item ContractType) (ContractType, error) {
    result := r.DB.Create(&item)
    if result.Error != nil {
        return ContractType{}, result.Error
    }
    return item, nil
}

func (r *Repository) GetAllContractTypes(limit, page int, filters []operators.FilterBlock) (items []ContractType, count int64) {
    _query := helper.QueryBuilder(ContractType{}, r.DB, filters)
    _query.Find(&items).Count(&count)
    _query.Scopes(pagination.NewPaginate(limit, page).PaginatedResult).Find(&items)
    return items, count
}

func (r *Repository) UpdateContractType(item ContractType) error {
    result := r.DB.Save(&item)
    return result.Error
}

func (r *Repository) DeleteContractType(item ContractType) error {
    result := r.DB.Delete(&item)
    return result.Error
}

func (r *Repository) FindContractTypeById(id int) (ContractType, error) {
    var item ContractType
    result := r.DB.First(&item, id)
    if result.Error != nil {
        return ContractType{}, errors.New("not found contracttype")
    }
    return item, nil
}
