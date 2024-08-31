package main

import "fmt"

type Component interface {
	Operation() string
}

type Leaf struct {
	name string
}

func (leaf *Leaf) Operation() string {
	return "leaf:" + leaf.name
}

type Composite struct {
	name     string
	children []Component
}

func (composite *Composite) Operation() string {
	result := "composite:" + composite.name

	for _, child := range composite.children {
		result += child.Operation() + "\n"
	}
	return result
}

func (composite *Composite) Add(component Component) {
	composite.children = append(composite.children, component)
}

func (composite *Composite) Remove(component Component) {
	for i, child := range composite.children {
		if child == component {
			composite.children = append(composite.children[:i], composite.children[i+1:]...)
			break
		}
	}
}

// func (composite *Composite) remove2(component Component){
// 	for i, child := range composite.children {
// 		if child == component {
// 			composite.children=append(composite.children[:i],composite.children[i+1:]...)
// 			break
// 	}
// }

func (composite *Composite) GetChild(index int) Component {
	return composite.children[index]
}

func main() {

	leaf1 := &Leaf{name: "leaf1"}
	leaf2 := &Leaf{name: "leaf2"}
	leaf3 := &Leaf{name: "leaf3"}

	composite1 := &Composite{name: "composite1"}
	composite2 := &Composite{name: "composite2"}

	composite1.Add(leaf1)

	composite2.Add(leaf2)

	composite2.Add(leaf3)

	composite1.Add(composite2)

	fmt.Println(composite1.Operation())

	fmt.Println(composite2.Operation())

	composite2.Remove(leaf3)

	fmt.Println(composite2.Operation())
}
