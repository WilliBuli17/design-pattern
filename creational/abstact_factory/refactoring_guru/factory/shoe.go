package factory

import "design-patterns/creational/abstact_factory/refactoring_guru/factory/constant"

type Shoe interface {
	Price() int64
	Type() constant.ShoeType
	IsNylon() bool
}
