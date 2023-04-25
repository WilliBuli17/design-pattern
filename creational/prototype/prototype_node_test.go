package prototype

import (
	"design-patterns/creational/prototype/refactoring_guru"
	"fmt"
	"testing"
)

func TestNodePrototype(t *testing.T) {
	file1 := &refactoring_guru.File{Name: "File1"}
	file2 := &refactoring_guru.File{Name: "File2"}
	file3 := &refactoring_guru.File{Name: "File3"}

	folder1 := &refactoring_guru.Folder{
		Child: []refactoring_guru.Node{file1},
		Name:  "Folder1",
	}
	folder2 := &refactoring_guru.Folder{
		Child: []refactoring_guru.Node{folder1, file2, file3},
		Name:  "Folder2",
	}

	fmt.Println("Printing hierarchy for Folder2")
	folder2.Print("   x")

	cloneFolder := folder2.Clone()
	fmt.Println("\n\nPrinting hierarchy for clone Folder2")
	cloneFolder.Print("   x")
}
