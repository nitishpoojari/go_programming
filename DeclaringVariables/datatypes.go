package main

import "fmt"

func main() {

	//declare and assign
	var final = 3

	//declare and initialize variable
	var intTotal int
	intTotal = 3

	//multiple variable assign
	var val1, val2 = 3, "string"

	//another way to assign value
	directAssign := "Hello World"

	//Constant variable

	const constVariable = 5

	//constVariable= 3    <--- will give error

	fmt.Println(final, intTotal, val1, val2, directAssign, constVariable)

}
