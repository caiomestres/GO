package main

import "fmt"

func main() {
	//MAIN TYPES
	//string
	//bool
	//int
	//int int8 int16 int32 int64
	// unsigned int (0+, no negative numbers) uint uint8 uint16 uint32 uint64 uintptr
	//byte - alias for uint8
	//rune - alias for int32
	//float32 float64
	//complex64 complex128

	var name = "Brad" // No need for var name string
	var age = 37
	//Shorthand
	name2, email := "Caio", "caiodev@gmail.com"

	const isCool = true
	size := 1.3
	fmt.Println(name, age, name2, isCool, email)
	fmt.Printf("%T\n", size) //https://golang.org/pkg/fmt/

}
