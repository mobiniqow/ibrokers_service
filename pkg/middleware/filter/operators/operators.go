package operators

type Operators interface {
	Equal() string
	LessThan() string
	GreaterThan() string
	LessThanEqual() string
	GreaterThanEqual() string
}
