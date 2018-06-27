package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func square(x *float64) {
	*x = *x * *x
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)

	squareX := 1.5
	square(&squareX)
	fmt.Println(squareX)

	xx := 1
	yy := 2
	swap(&xx, &yy)
	fmt.Println(xx, yy)
}

/*
Problems
1. To get the memory address of a variable use the & operator
2. Assign a value to a point by using the *{variable} = value
3. new(type) will create a new pointer
4. The value of x is 2.25
5. see swap
*/

func swap(x *int, y *int) {
	temp := new(int)
	*temp = *x
	*x = *y
	*y = *temp
}
