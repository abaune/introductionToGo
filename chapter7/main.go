package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}

func f2() (r int) {
	r = 1
	return
}

func f() (int, int) {
	return 5, 6
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	fmt.Println(f2())
	x, y := f()
	fmt.Println(x, y)
	fmt.Println(add(1, 2, 3))

	xs1 := []int{1, 2, 3}
	fmt.Println(add(xs1...))

	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))

	x1 := 0
	increment := func() int {
		x1++
		return x1
	}
	fmt.Println(increment())
	fmt.Println(increment())

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(factorial(3))

	defer second()
	first()

	fmt.Println(half(1))
	fmt.Println(half(2))
	fmt.Println(half(3))
	fmt.Println(half(4))

	testGreatestNum := []int{1, 2, 8, 1, 6, 9, 10, 1, 2, 3}
	fmt.Println(greatestNum(testGreatestNum...))

	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	fmt.Println(fib(10))

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

/* Problems
1. func sum(args ...int) int
2. see half
3. see greatestNum
4. see makeOddGenerator
5. see fib
6. Defer schedules a function to be called after the calling function completes. Panic is a function that calls a runtime error.
Recover stops the panic and returns the value that was passed to the panic call. Using a defer with a recover will allow you to recover
without the program stopping execution during the panic.
*/

func half(num int) (int, bool) {
	if num%2 == 0 {
		return num / 2, true
	}
	return 0, false
}

func greatestNum(args ...int) int {
	num := args[0]
	for _, v := range args {
		if v > num {
			num = v
		}
	}
	return num
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func fib(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
