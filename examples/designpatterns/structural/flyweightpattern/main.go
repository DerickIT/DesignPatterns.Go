package main

import "fmt"

type Shape interface {
	Draw()
}
type Circle struct {
	color string
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing a %s circle\n", c.color)
}

type ShapeFactory struct {
	shapes map[string]Shape
}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{
		shapes: make(map[string]Shape),
	}
}

func (f *ShapeFactory) GetCircle(color string) Shape {
	if shape, exists := f.shapes[color]; exists {
		return shape
	}
	circle := &Circle{color: color}
	f.shapes[color] = circle
	return circle
}

func main() {
	factory := NewShapeFactory()
	colors := []string{"red", "green", "blue", "red", "green", "blue"}

	for _, color := range colors {
		circle := factory.GetCircle(color)
		circle.Draw()
	}
	fmt.Printf("Total Circle Objects Created: %d\n", len(factory.shapes))
}
