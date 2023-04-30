package composite

import (
	"design-patterns/structural/composite/refactoring_guru/component_composite"
	"testing"
)

func TestComponentComposite(t *testing.T) {
	file1 := &component_composite.File{
		Name: "File_1",
	}
	file2 := &component_composite.File{
		Name: "File_2",
	}
	file3 := &component_composite.File{
		Name: "File_3",
	}

	folder1 := &component_composite.Folder{
		Name: "Folder_1",
	}
	folder2 := &component_composite.Folder{
		Name: "Folder_2",
	}

	folder1.Add(file1)

	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("rose")
}
