package filter

import (
	"reflect"
)

func GetAllowedFilters(model interface{}) []string {
	val := reflect.TypeOf(model)
	allowedFilters := []string{}

	// بررسی نوع و فیلدها
	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			name := field.Tag.Get("filter")
			if name != "" {
				allowedFilters = append(allowedFilters, name)
			}
		}
	}

	return allowedFilters
}
