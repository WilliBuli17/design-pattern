package factory

import "design-patterns/creational/abstact_factory/refactoring_guru/factory/constant"

type Shirt interface {
	Price() int64
	IsOversize() bool
	Material() constant.ShirtMaterial
}
