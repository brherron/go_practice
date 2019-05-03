package main

import "fmt"

var name = "Beau"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Beau"

	// Shorthand
	age := 26
	size := 1.3
	email := "beau@beauherrondev.com"

	// OR
	name, email := "Beau", "beau@beauherrondev.com"

	// Constant
	const isCool = true

	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n", email)
}
