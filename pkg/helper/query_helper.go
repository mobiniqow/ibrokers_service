package helper

import (
	"fmt"
	"gorm.io/gorm"
	"ibrokers_service/pkg/middleware/filter"
	"ibrokers_service/pkg/middleware/filter/operators"
	"slices"
)

func QueryBuilder(model interface{}, _query *gorm.DB, filters []operators.FilterBlock) (tx *gorm.DB) {
	allowedQuery := filter.GetAllowedFilters(model)

	for _, element := range filters {

		if slices.Contains(allowedQuery, element.Key) {
			_query = _query.Where(fmt.Sprintf("%s %s ?", element.Key, element.Operator), element.Value)
		}
	}
	return _query
}
