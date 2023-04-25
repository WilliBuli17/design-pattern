package refactoring_guru

import "fmt"

type File struct {
	Name string
}

func (f *File) Print(s string) {
	fmt.Println(s + f.Name)
}

func (f *File) Clone() Node {
	return &File{
		Name: f.Name + "_clone",
	}
}
