package main

import "fmt"

func main() {
	var name string
	var Surname string
	var age int
	fmt.Println("start...")
	fmt.Scanln(&name, &Surname, &age)
	fmt.Println("your name is", name, Surname, "and you are", age, "years old")
	fmt.Println("End...s")

}
