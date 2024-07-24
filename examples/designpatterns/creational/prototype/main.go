package main

import "fmt"

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

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, child := range f.children {
		if folder, ok := child.(*Folder); ok {
			folder.Print(indentation + " ")
		}
		if file, ok := child.(*File); ok {
			fmt.Println(indentation + " " + file.name)

		}
	}
}

func (f *Folder) AddChild(child Cloneable) {
	f.children = append(f.children, child)
}

func main() {
	root := &Folder{name: "root"}
	folder1 := &Folder{name: "Folder1"}
	folder2 := &Folder{name: "Folder2"}
	file1 := &File{name: "File1.txt"}
	file2 := &File{name: "File2.txt"}
	file3 := &File{name: "File3.txt"}

	folder1.AddChild(file1)
	folder2.AddChild(file2)
	folder2.AddChild(file3)
	root.AddChild(folder1)
	root.AddChild(folder2)
	fmt.Println("Original Folder Structure")
	root.Print("")

	cloneRoot := root.Clone().(*Folder)
	fmt.Println("\nCloned  Structure")
	cloneRoot.Print("")
	cloneFolder := cloneRoot.children[0].(*Folder)
	cloneFolder.AddChild(&File{name: "NewFile.txt"})

	fmt.Println("|nCloned Folder Structure")
	cloneRoot.Print("")

	fmt.Println("\nOriginal Folder Structure")
	root.Print("")

}
