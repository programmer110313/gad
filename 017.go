package main

import (
	"fmt"
	"reflect"
)

func main() {

	MyIntNumber := 5

	MyFloatNumber := float64(MyIntNumber)
	fmt.Println(MyFloatNumber)
	fmt.Println(reflect.TypeOf(MyFloatNumber))

}
