package main

import "fmt"

var globalX = "Global x"

func main() {
	var x string;
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	var x1 string = "hello"
	var y string = "world"
	fmt.Println(x1 == y)

	x2 := 5
	fmt.Println(x2)

	dogName := "Max"
	fmt.Println("The dog's name is ", dogName)
	f()
	exampleProgram()
	fahrenheitToCelsius()
	feetIntoMeters()
}

func f() {
	fmt.Println(globalX)
}

func exampleProgram() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

/* PROBLEMS
1. var x string; x := "string"
2. 6
3. const cannot be changed, var can
4. see below
5. see below
*/

func fahrenheitToCelsius() {
	fmt.Print("Enter degree in Farenheit:")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * 5/9

	fmt.Println(output)
}

func feetIntoMeters() {
	fmt.Print("Enter feet:")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 0.3048
	fmt.Println(output)
}
