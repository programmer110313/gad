package main

import (
	"fmt"
	"reflect"
)

func main() {
	do_military_service := true
	fmt.Println(do_military_service)
	fmt.Println(reflect.TypeOf(do_military_service))
}
