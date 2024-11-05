package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number float64 = 40.5
	fmt.Println(number)
	fmt.Println(reflect.TypeOf(number))
}
