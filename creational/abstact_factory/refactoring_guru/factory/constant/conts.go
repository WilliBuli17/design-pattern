package constant

type ShirtMaterial string
type ShoeMaterial string
type ShirtType string
type ShoeType string

const (
	Sintetis   ShirtMaterial = "Sintetis"
	Microfiber ShirtMaterial = "Microfiber"
	Katun      ShirtMaterial = "Katun"
	Calico     ShirtMaterial = "Calico"
	Spandex    ShirtMaterial = "Spandex"
)

const (
	Oversize    ShirtType = "Oversize"
	Undershirts ShirtType = "Undershirts"
)

const (
	Nylon ShoeMaterial = "Nylon"
	Lycra ShoeMaterial = "Lycra"
)
const (
	Climbing   ShoeType = "Climbing Shoes"
	Sneakers   ShoeType = "Sneakers / Trainers"
	Soccer     ShoeType = "Soccer Shoes"
	Basketball ShoeType = "Basketball Shoes"
	Golf       ShoeType = "Golf Shoes"
)
