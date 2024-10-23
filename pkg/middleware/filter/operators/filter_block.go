package operators

import "fmt"

type FilterBlock struct {
	Key      string
	Operator string
	Value    string
}

func (c *FilterBlock) Command() string {
	command := fmt.Sprintf("%s%s%s", c.Key, c.Operator, c.Value)
	return command
}
