package main

import "fmt"

func main() {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")
	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()
	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShirtDetails(nikeShirt)
	printShoeDetails(nikeShoe)

	printShirtDetails(adidasShirt)
	printShoeDetails(adidasShoe)
	// adidasFactory := Adidas{}
	// nikeFactory := Nike{}
	// shoe1 := adidasFactory.makeShoe()
	// shoe1.setLogo("Adidas")
	// shoe1.setSize(42)
	// fmt.Println(shoe1.getLogo())
	// fmt.Println(shoe1.getSize())
	// shirt1 := adidasFactory.makeShirt()
	// shirt1.setLogo("Adidas")
	// shirt1.setSize(42)
	// fmt.Println(shirt1.getLogo())
	// fmt.Println(shirt1.getSize())
	// shoe2 := nikeFactory.makeShoe()
	// shoe2.setLogo("Nike")
	// shoe2.setSize(41)
	// fmt.Println(shoe2.getLogo())
	// fmt.Println(shoe2.getSize())
	// shirt2 := nikeFactory.makeShirt()
	// shirt2.setLogo("Nike")
	// shirt2.setSize(41)
	// fmt.Println(shirt2.getLogo())
	// fmt.Println(shirt2.getSize())
}

func printShirtDetails(s IShirt) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d\n", s.getSize())
	fmt.Println()
}

func printShoeDetails(s IShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d\n", s.getSize())
	fmt.Println()
}
