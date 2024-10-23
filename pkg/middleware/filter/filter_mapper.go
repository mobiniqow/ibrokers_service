package filter

import (
	"ibrokers_service/pkg/middleware/filter/operators"
	"strings"
)

type Mapper struct {
	Operators operators.Operators
}

func (c *Mapper) Convert(filters []string) []operators.FilterBlock {
	var _operators []operators.FilterBlock
	for _, element := range filters {
		if strings.Contains(element, "=") {
			query := strings.Split(element, "=")
			if len(query) == 2 {
				_key, operator, isMulti := c.SplitCommand(query[0])
				if isMulti {
					operator := c.CheckOperator(operator)
					_operators = append(_operators, operators.FilterBlock{Key: _key, Operator: operator, Value: query[1]})
				} else {
					_operators = append(_operators, operators.FilterBlock{Key: query[0], Operator: c.Operators.Equal(), Value: query[1]})
				}
			}
		}
	}

	return _operators
}

func (c *Mapper) SplitCommand(key string) (_key string, operator string, isMulti bool) {
	if strings.Contains(key, "_") {
		parts := strings.Split(key, "_")
		_key = parts[0]
		operator = parts[1]
		isMulti = true
	} else {
		_key = key
		operator = ""
		isMulti = false
	}
	return _key, operator, isMulti
}

func (c *Mapper) CheckOperator(operator string) string {
	switch operator {
	case "gt":
		return c.Operators.GreaterThan()
	case "lt":
		return c.Operators.LessThan()
	case "lte":
		return c.Operators.LessThanEqual()
	case "gte":
		return c.Operators.GreaterThanEqual()
	default:
		return c.Operators.Equal()
	}
}
