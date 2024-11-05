package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Azam"
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))
	full_name := "Azam mashhadi"
	fmt.Println(full_name)
	fmt.Println(reflect.TypeOf(full_name))
}
