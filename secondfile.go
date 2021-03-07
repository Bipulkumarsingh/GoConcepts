package main

import "fmt"

func second() {
	// Boolean variable
	var h bool = true
	fmt.Printf("%v, %T\n", h, h)
	i := 1 == 1 // double = is used to check both value are same or not
	fmt.Printf("%v, %T\n", i, i)

	// unlike signed integer we also have unsigned integer
	// we don't have 64 bit unsigned integer

	var j uint16 = 42
	fmt.Printf("%v, %T\n", j, j)
	var k int = 10
	var l int8 = 3
	// fmt.Println(k + l) // we are not allowed to add two integer of different type
	fmt.Println(k + int(l)) // we need to convert one of them to add it.
	// Bit operators
	m := 56
	fmt.Println(k & m)  // And operator
	fmt.Println(k | m)  // OR operator
	fmt.Println(k ^ m)  // Exclusive OR operaton
	fmt.Println(k &^ m) // And Not operator

	n := "This is a string"
	o := []byte(n)
	fmt.Printf("%v, %T\n", o, o) // print ASCII value of each letter. and it type of uint8

	// string represent utf-8 character
	// Rune represent utf-32 character
	var p rune = 'a'
	fmt.Printf("%v, %T\n", p, p) // it will return 97, type int32 Rune is just like alias of int32

}
