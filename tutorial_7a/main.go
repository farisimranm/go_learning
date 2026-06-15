package main

import "fmt"

func main() {
	// pointers + functions

	// 1. passing a list as argument
	floatList := [5]float32{1, 2, 3, 4, 5}
	resultList := square(floatList)

	fmt.Println("floatList:", floatList)
	fmt.Println("resultList:", resultList)

	// 2. passing the pointer of a list as argument
	floatList2 := [5]float32{5, 4, 3, 2, 1}

	fmt.Println("\nfloatList2:", floatList2)

	squareViaPointer(&floatList2)

	fmt.Println("updated floatList2:", floatList2)

	// 3. passing the pointer of a var as argument
	str := "deez"

	fmt.Println("\nstr:", str)

	updateStringViaPointer(&str)

	fmt.Println("updated str:", str)
}

// takes a list of floats and return a list of squared elements
func square(floatList [5]float32) [5]float32 {
	var resultList [5]float32

	for i, v := range floatList {
		resultList[i] = v * v
	}

	return resultList
}

// takes a pointer of a list of floats and square all elements
// note: void method. will modify the actual input list
func squareViaPointer(floatList *[5]float32) {
	for i, v := range floatList {
		floatList[i] = v * v
	}
}

func updateStringViaPointer(str *string) {
	*str = "nuts"
}
