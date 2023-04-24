package other

type FurnitureFactory interface {
	CreateChair() Chair
	CreateTable() Table
	CreateSofa() Sofa
}
