package other

type SofaStyle string

const (
	EuropeanStyle SofaStyle = "European"
	AmericanStyle SofaStyle = "American"
)

type Sofa interface {
	IPrice
	Style() SofaStyle
}
