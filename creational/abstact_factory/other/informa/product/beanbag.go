package product

type BeanBag struct {
	ProductPrice  int64
	SoftnessLevel int
}

func (b *BeanBag) Price() int64 {
	return b.ProductPrice
}

func (b *BeanBag) IsErgonomic() bool {
	return b.SoftnessLevel > 10
}
