package main

import "fmt"

type Visitor interface {
	VisitCircle(c *Circle)
	VisitRectangle(e *Rectangle)
}

type Shape interface {
	Accept(v Visitor)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitRectangle(r)
}

type AreaCalculator struct{}

func (ac *AreaCalculator) VisitCircle(c *Circle) {
	fmt.Printf("Area of Circle with radius %.2f is %.2f\n", c.Radius, 3.14*c.Radius*c.Radius)
}

func (ac *AreaCalculator) VisitRectangle(r *Rectangle) {
	fmt.Printf("Area of Rectangle with width %.2f and height %.2f is %.2f\n", r.Width, r.Height, r.Width*r.Height)
}

type PerimeterCalculator struct{}

func (pc *PerimeterCalculator) VisitCircle(c *Circle) {
	fmt.Printf("Perimeter of Circle with radius %.2f is %.2f\n", c.Radius, 2*3.14*c.Radius)
}

func (pc *PerimeterCalculator) VisitRectangle(r *Rectangle) {
	fmt.Printf("Perimeter of Rectangle with width %.2f and height %.2f is %.2f\n", r.Width, r.Height, 2*(r.Width+r.Height))
}

func main() {
	shapts := []Shape{
		&Circle{Radius: 2},
		&Rectangle{Width: 3, Height: 4},
	}
	areaCalc := &AreaCalculator{}
	perimeterCalc := &PerimeterCalculator{}

	for _, shape := range shapts {
		shape.Accept(areaCalc)
		shape.Accept(perimeterCalc)
		fmt.Println()
	}
}
