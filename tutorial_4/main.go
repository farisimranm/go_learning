package main

import (
	"fmt"
	"time"
)

func main() {
	arrayLesson()
	sliceLesson()
	mapLesson()
	loopLesson()
}

func arrayLesson()  {
	// arrays are fixed-length
	// 1. unintialized array
	var intArray [3]int32

	fmt.Println("intArray ->", intArray)

	// 2. intialized array
	var intArray2 = [5]int32{5, 4, 3, 2, 1}

	fmt.Println("\nintArray2 ->", intArray2)
	fmt.Println("intArray2[3] ->", intArray2[3])
	fmt.Println("intArray2[1:4] ->", intArray2[1:4])

	// 3. intialized array (different syntax)
	intArray3 := [...]int32{5, 4, 3, 2, 1}

	fmt.Println("\nintArray3 ->", intArray3)

	// iterating an array
	for i, v := range intArray3 {
		fmt.Printf("i:%v, v:%v\n", i, v)
	}
}

func sliceLesson() {
	// slices are dynamic-length array
	// 1. initialized slice
	stringSlice := []string{"hello", "deez", "nuts"}

	fmt.Println("\nstringSlice ->", stringSlice)
	fmt.Println("len(stringSlice) ->", len(stringSlice))
	fmt.Println("cap(stringSlice) ->", cap(stringSlice))

	// 2. slice via append(array, newValue)
	stringSlice2 := append(stringSlice, "gottem")

	fmt.Println("\nstringSlice2 ->", stringSlice2)
	fmt.Println("len(stringSlice2) ->", len(stringSlice2))
	fmt.Println("cap(stringSlice2) ->", cap(stringSlice2))
	fmt.Println("stringSlice2[2] ->", stringSlice2[2])
	fmt.Println("stringSlice2[1:] ->", stringSlice2[1:])

	// 3. slice via make(type, len, cap)
	stringSlice3 := make([]string, 2, 6)
	stringSlice3[0], stringSlice3[1] = "learning", "go"

	fmt.Println("\nstringSlice3 ->", stringSlice3)
	fmt.Println("len(stringSlice3) ->", len(stringSlice3))
	fmt.Println("cap(stringSlice3) ->", cap(stringSlice3))

	// 4. slice optimization: slice with preallocated capacity is faster than without
	var n int = 1000000 					// one million
	var testSlice = []int{}					// non-preallocated capacity
	var testSlice2 = make([]int, 0, n)		// preallocated capacity

	fmt.Println("\nComparing preallocated and non-preallocated slices performance")
	fmt.Printf("Total time without preallocation: %v\n", timeToLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeToLoop(testSlice2, n))
}

func timeToLoop(slice []int, count int) time.Duration {
	var now = time.Now()
	for len(slice) < count {
		slice = append(slice, 0)
	}
	return time.Since(now)
}

func mapLesson() {
	// map
	// 1. uninitialized map
	var stringUintMap = make(map[string]uint)
	fmt.Println("\nstringUintMap ->", stringUintMap)

	// 2. initialized map
	stringUintMap2 := map[string]uint{"arthur":47, "john":35}
	fmt.Println("\nstringUintMap2 ->", stringUintMap2)
	fmt.Println("stringUintMap2[\"arthur\"] ->", stringUintMap2["arthur"])

	// NOTE: map will return default value if key does not exist
	fmt.Println("stringUintMap2[\"dutch\"] ->", stringUintMap2["dutch"])

	// map will return value and boolean. catch the boolean to check if key exists
	mapVal, isKeyFound := stringUintMap2["dutch"]
	if isKeyFound {
		fmt.Println("key \"dutch\" found with value: ", mapVal)
	} else {
		fmt.Println("key \"dutch\" does NOT found")
	}

	// deleting a key from a map
	delete(stringUintMap2, "john")
	fmt.Println("\nstringUintMap2 ->", stringUintMap2)

	// 3. iterating a map
	stringUintMap3 := map[string]uint{"ronaldo": 973, "mbappe": 425, "kane": 500}

	fmt.Println("\nstringUintMap3 ->", stringUintMap3)

	// order is not preserved
	for playerName, totalGoals := range stringUintMap3 {
		fmt.Printf("%v has scored %v goals\n", playerName, totalGoals)
	}
}

func loopLesson() {
	fmt.Println()

	// 1. loop using condition
	fmt.Println()

	var i int = 0
	for i < 7 {
		fmt.Println("i:", i)
		i++
	}

	// 2. loop using declared index
	fmt.Println()

	var count int = 7
	for j := 0; j < count; j++ {
		fmt.Println("j:", j)
	}

	// 2. loop without condition. use break
	fmt.Println()

	var k int = 0
	for {
		if k >= 7 {
			break
		}

		fmt.Println("k:", k)
		k++
	}
}