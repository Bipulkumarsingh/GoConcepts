package main

import (
	"fmt"
	"reflect"
)

// struct type 
type Doctor struct {
	number int
	actorName string
	companions []string
}
// if we are working with web application and getting input from user, we want some specific field not to be blank and also
// want to fix the length , We can use use tags 
type Animal struct {
	Name string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly bool
}

func fourth(){
	// map is just like dictionary in python.
	statePopulations := map[string]int{
		"California": 34929,
		"Texas": 34343,
		"Florida": 43423,
		"New York": 9394,
		"Pennsylvania": 343224,
		"Ohio": 32423423,
	}
	statePopulations["Geogia"] = 334343
	
	//delete value from map
	delete(statePopulations, "Geogia")
	fmt.Println(statePopulations)

	// if we search for "Geogia" in map like statePopulations["Geogia"]
	// it will return 0 instead of error, we can check is value is zero or key not available.
	pop, ok := statePopulations["Geogia"] // here ok variable store true or false for key
	// if we want to check only key is present or not we can do like this.
	_, is_present := statePopulations["Geogia"]
	fmt.Println(pop, ok, is_present)
	// length of map
	fmt.Println(len(statePopulations))

	sp := statePopulations
	delete(sp, "Ohio") // it will impact statePopulations, 
	// now create a object of struct 
	aDoctor := Doctor{
		number: 3,
		actorName: "Jone Pertwee",
		companions: []string{
			"LIz Shaw",
			"Jo Grant",
		},
	}
	fmt.Println(aDoctor)

	// Anonymous struct, we can't use this anywhere because this don't have any name.
	aDoctor1 := struct{name string}{name: "John Pertwee",}
	fmt.Println(aDoctor1)
	anotherDocter := &aDoctor1
	anotherDocter.name = "tom Baker" // It will change value of both aDoctor1 and anotherDocter

	// Go Doesn't support traditional OOPS concept so GO don't have Inheritance.
	// Go support composition called embedeging
	b := Bird{} //istance 
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
	// we can get tag directly from stuct object, we need to use reflect for this.
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag) 
}