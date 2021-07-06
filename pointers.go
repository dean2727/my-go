package main

import "fmt"

// Go is pass-by-value
// syntax for pointers and references is very similar to c++

// pass-by-reference. argument is a pointer, and then dereference it in the function to change the value
func addHundred (numPtr *int) {
	*numPtr += 100
}

func main() {
	x := "My very first address"
	fmt.Println(&x) // Prints 0x414020

	var pointerForInt *int 
	minutes := 525600
	pointerForInt = &minutes
	// implicit pointer declaration
	otherPointer := &minutes
	fmt.Println(pointerForInt) // Prints 0xc000018038

	// dereferencing/indirecting
	lyrics := "Moments so dear" 
	pointerForStr := &lyrics
	*pointerForStr = "Journeys to plan" 
	fmt.Println(lyrics) // Prints: Journeys to plan
}