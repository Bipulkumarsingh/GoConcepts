package main

import (
	"fmt"
)

func sixth(){
	var a int = 42 
	var b *int = &a // *int is used to declare pointer to datatype
	fmt.Println(a, b, *b) // Normally printing b print address of a , we need to print *b to show value.
	// *b means dereference 
	// changing value of a or b changes reflected in both variables. 	

	c := [3]int{1,2,3}
	d := &c[0]
	e := &c[1]
	fmt.Printf("%v %p %p\n", c,d,e)

	// Functions in GO
	SayMessage("Hello Go")
	for i :=0;i<5;i++{
		SayM("Hello", i)
	}
	greet := "Hello"
	name := "stacey"
	sayGreet(&greet, &name)
	fmt.Println(name)

	value, err := sum("hello", 1,2,3,4,5)
	if err != nil{
		fmt.Println("error occur")
		return
	}
	fmt.Println(value)

	// Anonymous function
	func() {
		fmt.Println("Hello Go!")
	}()
}

func SayMessage(msg string){
	fmt.Println(msg)
}

func SayM(msg string, idx int){
	fmt.Println(msg)
	fmt.Println("The value of the index is", idx)
}

// func sayGreeting(greeting string, name string)
func sayGreeting(greeting, name string){  // if we have same type of parameters
	fmt.Println(greeting, name)
}
// Above function are pass by value

// Pass by reference
func sayGreet(greeting, name *string){
	fmt.Println(*greeting, *name)
	*name = "Ted" // chaning name here also change name from original place
	fmt.Println(name)
}

func sum(msg string, values ...int) (int, error){  // three dots takes all value and wrap it into slice, Put varatic parameter at the end
	fmt.Println(values)
	result := 0
	for _,v := range values{
		result +=v
	}
	if result == 0{
		return 0, fmt.Errorf("Don't have any value")
	}
	fmt.Println("The sum is", result)
	return result, nil
}
