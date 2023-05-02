package refactoring_guru

type Server interface {
	HandleRequest(string, string) (int, string)
}
