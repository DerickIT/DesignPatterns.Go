package main

type Cloneable interface {
	Clone() Cloneable
}

type File struct {
	name string
}

func (f *File) Clone() Cloneable {
	return &File{name: f.name}
}

type Folder struct {
	name     string
	children []Cloneable
}

func (f *Folder) Clone() Cloneable {
	cloneFolder := &Folder{name: f.name}
	var tempChildren []Cloneable
	for _, child := range f.children {
		tempChildren = append(tempChildren, child.Clone())
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}
