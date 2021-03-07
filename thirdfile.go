package main

import (
	"fmt"
)

const myConst int16 = 27
const q = iota // iota start with 0

const (
	q0 = iota // value 0
	q1 = iota // value 1
	q2 = iota // value 2
	q3 = iota // value 3
)

// iota is scoped to constant block, above constant
// we don't need to assign iota if we are declaring constant in pattern
const (
	_  = iota // you can assign 0 value to this underscore
	q4        // value 1
	q5        // value 2
	q6        // value 3
	q7        // value 4
)

func third() {
	// const value, we can create constant with any primitive type of GO
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst) // It will show inner declared constant

	// if variable and constant are of same type we can perform operation on that.
	fmt.Printf("%v, %T\n", q, q)

	// Array
	grades := [...]int{97, 85, 93} // if we are assinging the value at the time of declaration we can also provide three
	// dots instead of size
	fmt.Printf("Grades: %v, %v, %v\n", grades[0], grades[1], grades[2])

	// Declare an array but not assigned value
	var students [3]string
	fmt.Printf("Student %v", students)
	students[0] = "Lisa"
	students[1]= "Arnold"
	students[2] = "Ahmed"
	fmt.Println(students)
	// Matrix of array
	var identityMatrix [3][3]int = [3][3]int{ [3]int{1,2,3}, [3]int{4,5,6}, [3]int{7,8,9}}
	fmt.Println(identityMatrix)

	// we can write this matrix in more cleaner way
	var indentitym [3][3]int
	indentitym[0] = [3]int{0,1,2}
	indentitym[1] = [3]int{3,4,5}
	indentitym[2] = [3]int{6,7,8}
	fmt.Println(indentitym)

	// if we pass value of array variable to another variable it will create copy of it
	// changing value of second variable will not affect first one. 

	// if we are passing array to function, try to pass reference of array because passing by value create a new copy 
	// if array is small it will not affect anyting but if array contain millions of value then it's a big deal
	a := [...]int{1,2,3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
	c := &a // Now a is passing it's addresss instead of whole value.
	c[1] = 7 // it will also change the value of a
	// Slices are approx same as array
	d := []int{1,2,4}
	e := d
	e[1] = 9
	// slice have default address reference if we assign slice to different variable
	// array need to use & sign
	fmt.Printf("Length: %v\n", len(d))
	fmt.Printf("Capacity: %v\n", cap(d)) // capacity is extra feature which is not available in array
	
	// Other way of creating slices
	g := []int{1,2,3,4,5,6}
	h := g[:] // slice of all elements
	i := g[3:] // slice from 4th element to end
	j := g[:6] // slice first 6 elements
	k := g[3:6] //slice the 4th, 5th and 6th element
	fmt.Println(g, h, i, j, k)
	// create blank slice
	l := []int{}
	l = append(l, 1)
	l = append(l, 3,4,5,5)

	// if we have another slice and want to insert into other slice , append method not allow to append it directly
	// we need to spread each value because append only accept integer value for integer slice.
	l = append(l, []int{8,9,5}...) // three dots will spread the value into integer.
	// slice is just like stack and append is used to push element into stack
	
	// pop element 
	l = l[1:] // it will remove first element
	l = l[:len(l)-1] // it will remove last element
}

