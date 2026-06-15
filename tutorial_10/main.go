package main

import "fmt"

func main() {
	// 1. without using generic
	var slice1 = []int{1, 2, 3}
	fmt.Println("sumIntSlice:", sumIntSlice(slice1))

	var slice2 = []float32{1.1, 2.2, 3.3}
	fmt.Println("sumFloat32Slice:", sumFloat32Slice(slice2))

	var slice3 = []float64{1.11, 2.22, 3.33}
	fmt.Println("sumFloat64Slice:", sumFloat64Slice(slice3))

	// 2. using generic
	var slice4 = []int{1, 2, 3}
	var slice5 = []float32{1.1, 2.2, 3.3}
	var slice6 = []float64{1.11, 2.22, 3.33}
	fmt.Println("\nsumSlice:", sumSlice(slice4))
	fmt.Println("sumSlice:", sumSlice(slice5))
	fmt.Println("sumSlice:", sumSlice(slice6))

	// 3. using generic "any"
	var emptySlice []int
	var emptySlice2 []string
	var emptySlice3 = make([]float64, 0)

	fmt.Println("\nisEmpty(emptySlice):", isEmpty[int](emptySlice))		// putting explicit type
	fmt.Println("isEmpty(emptySlice2):", isEmpty(emptySlice2))			// omitting type
	fmt.Println("isEmpty(emptySlice3):", isEmpty(emptySlice3))
}

func sumIntSlice(inputSlice []int) int {
	var sum int = 0
	for _, v := range inputSlice {
		sum += v
	}
	return sum
}

func sumFloat32Slice(inputSlice []float32) float32 {
	var sum float32 = 0
	for _, v := range inputSlice {
		sum += v
	}
	return sum
}

func sumFloat64Slice(inputSlice []float64) float64 {
	var sum float64 = 0
	for _, v := range inputSlice {
		sum += v
	}
	return sum
}

// this function accepts 3 generic types as argument and output
func sumSlice[T int | float32 | float64](inputSlice []T) T {
	var sum T = 0
	for _, v := range inputSlice {
		sum += v
	}
	return sum
}

// this function accepts any type
func isEmpty[T any](value []T) bool {
	return len(value) == 0
}
