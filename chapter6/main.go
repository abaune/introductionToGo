package main

import "fmt"

func main() {
	// var x [5]float64
	// x[0] = 98
	// x[1] = 93
	// x[2] = 77
	// x[3] = 82
	// x[4] = 83
	x := [5]float64{98, 93, 77, 82, 83}

	var total float64 = 0
	// for i := 0; i < len(x); i++ {
	// 	total += x[i]
	// }

	for _, value := range x {
		total += value
	}

	fmt.Println(total / float64(len(x)))

	sliceTest()
	mapTest()
	getMin()
}

func sliceTest() {
	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[0:5] // can omit either high or low
	fmt.Println(x)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := make([]int, 2)
	copy(slice3, slice1)
	fmt.Println(slice1, slice3)
}

func mapTest() {
	x := make(map[string]int)

	x["key"] = 10
	fmt.Println(x["key"])

	x2 := make(map[int]int)

	x2[1] = 10
	fmt.Println(x2[1])

	exampleMapProgram()
}

func exampleMapProgram() {
	// elements := make(map[string]string)
	// elements["H"] = "Hydrogen"
	// elements["He"] = "Helium"
	// elements["Li"] = "Lithium"
	// elements["Be"] = "Beryllium"
	// elements["B"] = "Boron"
	// elements["C"] = "Carbon"
	// elements["N"] = "Nitrogen"
	// elements["O"] = "Oxygen"
	// elements["F"] = "Fluorine"
	// elements["Ne"] = "Neon"

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}

/* PROBLEMS
1. arr[5] Gives the 4th element in an array or slice
2. 3
3. [c, d, e]
4. See below
*/

func getMin() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	min := x[0]

	for i := 0; i < len(x); i++ {
		if x[i] < min {
			min = x[i]
		}
	}

	fmt.Println(min)
}
