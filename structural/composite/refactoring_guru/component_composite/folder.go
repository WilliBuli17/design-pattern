package component_composite

import (
	"design-patterns/structural/composite/refactoring_guru"
	"fmt"
)

type Folder struct {
	components []refactoring_guru.Component
	Name       string
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Search recursively for keyword %s in folder %s\n", keyword, f.Name)

	for _, component := range f.components {
		component.Search(keyword)
	}
}

func (f *Folder) Add(c refactoring_guru.Component) {
	f.components = append(f.components, c)
}
