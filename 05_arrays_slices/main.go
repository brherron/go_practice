package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string
	fruitArr := [2]string{"Apple", "Orange"}

	// Assign vals
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(len(fruitArr))
}
