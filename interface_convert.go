package main

import (
	"fmt"
)

type Host struct {
	Name string
}

func maininterface() {
	// You have to explicitly convert your interface{} to a slice of the expected type to achieve it.
	Map := make(map[string]interface{})
	Map["hosts"] = []Host{Host{"test.com"}, Host{"test2.com"}}

	hm := Map["hosts"].([]Host)
	fmt.Println(hm[0])
}
