package operators

type GormOperator struct {
}

func (c GormOperator) Equal() string {
	return "="
}

func (c GormOperator) LessThan() string {
	return "<"
}

func (c GormOperator) GreaterThan() string {
	return ">"
}

func (c GormOperator) LessThanEqual() string {
	return "<="
}

func (c GormOperator) GreaterThanEqual() string {
	return ">="
}
