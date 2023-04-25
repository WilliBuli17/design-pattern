package refactoring_guru

type Node interface {
	Print(string)
	Clone() Node
}
