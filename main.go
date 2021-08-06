package main

import "fmt"

func main() {
	// Factory generation
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")
	// Shoe generation
	aNikeShoe := nikeFactory.makeShoe()
	// Short generation
	aNikeShort := nikeFactory.makeShort()
	// Shoe generation
	anAdidasShoe := adidasFactory.makeShoe()
	// Short generation
	anAdidasShort := adidasFactory.makeShort()

	printShoeDetails(aNikeShoe)
	printShortDetails(aNikeShort)
	printShoeDetails(anAdidasShoe)
	printShortDetails(anAdidasShort)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}
func printShortDetails(s iShort) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}
