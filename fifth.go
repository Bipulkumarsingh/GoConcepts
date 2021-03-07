package main

import(
	"fmt"
	"strconv"
)

func fifth(){
	if true{
		fmt.Println("working just fine")
	}
	simpleMap := map[string]int{
		"first":1,
		"second":2,

	}
	if pop, ok := simpleMap["first"]; ok{
		fmt.Println("first value is present: "+strconv.Itoa(pop))
	}

	// switch 
	i := 10 
	switch 2 {
	case 1:
		fmt.Println("This is one")
		// break we can use break according to requirement
		// fallthrough - it is used to execute next statement no matter case is valid or not
	case 2:
		fmt.Println("THis is two")
	case 3:
		fmt.Println("This is three")
	case 4, 5, 6:
		fmt.Println("We can check multiple statement in one case")
	case i:
		fmt.Println("we can also check variable value")
	// case i <= 20:
	// 	fmt.Println("we can also apply conditions on check")
	default:
		fmt.Print("NOne of the above case matched")
	}
	var j interface{} = 1
	switch j.(type) {
	case int:
		fmt.Println("j is an int")
	case float64:
		fmt.Println("j is a float")
	case string:
		fmt.Println("j is a string")
	default:
		fmt.Println("J is another type")
	}

	// for loop
	for i := 0; i<5; i++{
		fmt.Println(i)
	}
	k := 0
	for ; k<5; k++{
		fmt.Println(k)
	}
	// for ;k<5;{ // we can also do that but it stuck in infinite loop
	// 	fmt.Println(k)
	// k++ we can also increase value inside for loop
	// }
	// for k<5{  this is also valid statement of for loop
	// 	k++
	// }

	for {
		fmt.Println(k)
		k++
		if k == 2{
			continue
		}
		if k == 5 {
			break 
		}	
	}
	s := "This is string"
	for y,z := range s {
		fmt.Println(y,string(z))
	}
	// Defer is the special keyword which execute function or statement at the end of the function
	// but before return function.
	defer fmt.Println("Print at the end of the main function")
	fmt.Println("Print before the defer")
	// Defer execute in LIFO order means last in first out
	defer fmt.Println("This will print before above defer.")

	// Panic is just like try It will execute when error occurs.
	fmt.Println("start")
	panic("something bad happens")
}