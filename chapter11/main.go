package main

import "fmt"
import m "introductionToGo/chapter11/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)

	test := []float64{4, 2, 5, 2, 6, 1, 3}
	fmt.Println(m.Min(test)) // should be 1
	fmt.Println(m.Max(test)) // should be 6
}

/**
Problems
1. Packages are separate programs that can be used for proper code reuse. They also help during compilation.
2. Capital letters are public and lower case are private.
3. A package alias is another name for an imported package. For example the math package could be imported as `import m "path/to/package"`.
4. see math/math.go
5. see math/math.go
*/
