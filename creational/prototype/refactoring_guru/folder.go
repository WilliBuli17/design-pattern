package refactoring_guru

import "fmt"

type Folder struct {
	Child []Node
	Name  string
}

func (f *Folder) Print(s string) {
	fmt.Println(s + f.Name)

	for _, node := range f.Child {
		node.Print(s + s)
	}
}

func (f *Folder) Clone() Node {
	cloneFolder := &Folder{
		Name: f.Name + "_clone",
	}

	var tempChild []Node
	for _, node := range f.Child {
		clone := node.Clone()
		tempChild = append(tempChild, clone)
	}

	cloneFolder.Child = tempChild
	return cloneFolder
}
