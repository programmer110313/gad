package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "fatema"
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))
}
