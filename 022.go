package main

import (
	"fmt"
)

func main() {
	MyNumbers := [5]int{18, 12, 15, 17, 20}
	for _, item := range MyNumbers {
		fmt.Println(item)
	}

}
