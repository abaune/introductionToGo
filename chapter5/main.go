package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " even")
		} else {
			fmt.Println(i, " odd")
		}
	}

	for i := 1; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		case 6:
			fmt.Println("Six")
		case 7:
			fmt.Println("Seven")
		case 8:
			fmt.Println("Eight")
		case 9:
			fmt.Println("Nine")
		default:
			fmt.Println("Unknown Number")
		}
	}

	fmt.Println()
	printNumDivBy3()
	fmt.Println()
	fizzBuzz()
}

/* Problems
1. Small
2. See below
3. See below
*/

func printNumDivBy3() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		value := ""
		if i%3 == 0 {
			value = "fizz"
		}

		if i%5 == 0 {
			value += "buzz"
		}

		if value == "" {
			fmt.Println(i)
		} else {
			fmt.Println(value)
		}
	}
}
