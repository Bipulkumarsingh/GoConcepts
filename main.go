package main

import (
	"fmt"
	"strconv"
)

/*
	There are three scope of variable and use cases
	1. If we create a variable inside scope, we can declare in camelCase and is visible or used within that scope,
		means I have a variable name 'a' in main , we can use 'a' within main.
	2. Local package level variable, we can declare variable just after import or outside function with Camelcase format,
		these variable is visible to all files of same package.
	3. Global package level variable, we can declare variable for other package as well, these type of variable starts
		with uppercase letter
	Rule for acronym
		If we are declaring variable with acronym we need to write acronym in uppercase like theHTTP(http is abrivation of
		hyper text transfer protocol)

*/

// We can also declare variable at package level but if we declare variable at
// package level we can't use ':='
var f1 int = 45

// At package level if we declare multiple variable then we can combine them with same var keyword

var (
	actorName  string = "Elisabeth Sladen"
	companion  string = "Sarah Jane Smith"
	doctorName int    = 3
	season     int    = 1
)

// we can define multiple var block to organise variable of same context
var (
	counter int = 0
)

func main() {
	//create variable
	var a int
	a = 43
	a = 56
	fmt.Println(a)
	// we can also create variable and assign value at same time
	var b int = 67
	fmt.Println(b)
	// Go compilar figure out what datatype we need to have when we assign value
	// so we also not need to pass var and data type
	c := 78
	fmt.Printf("%v, %T\n", c, c)
	// if we want to assign value as float
	var d float32 = 27
	fmt.Println(d)
	e := 99.
	fmt.Printf("%v, %T\n", e, e) // this is default float64
	fmt.Println(f1)

	// there are two scope of variable one is package and another is local scope
	fmt.Println(season) // Value is 1
	season := 56
	fmt.Println(season) // Now value is 56
	var f int = 23
	//if we want to convert int into float, we can explicitly convert it
	var f2 float32
	f2 = float32(f)
	fmt.Printf("%v, %T\n", f2, f2)
	//If we try to convert int into string it will print "*", string is treated as array of bytes in Go
	var g string
	g = string(f) // only astrics will be printed, we need strconv package to convert int into string
	g = strconv.Itoa(f)
	fmt.Printf("%v, %T\n", g, g)

	//Functions are defined in different file
	second()
}
