package main

import "fmt"

func main()  {
	// pointers store memory address of variables
	// 1. initializing a pointer with existing variable
	var i int32 = 123
	var p *int32 = &i

	fmt.Println("i:", i)
	fmt.Println("p:", p)
	fmt.Println("value of p:", *p)

	// updating i
	i = 321

	fmt.Println("\nupdated i:", i)
	fmt.Println("p:", p)
	fmt.Println("value of p:", *p)
	//  ↳ the value of iPointer will reflect the new value. basically like Java side effect

	// 2. initializing a pointer without existing variable
	var p2 *int32 = new(int32)

	fmt.Println("\np2:", p2)
	fmt.Println("value of p2:", *p2)

	// update the value of pointer
	*p2 = 67
	fmt.Println("updated value of p2:", *p2)

	// 3. copying slices actually creates a pointer to the original slice
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice

	fmt.Println("\nslice:", slice)
	fmt.Println("sliceCopy:", sliceCopy)

	// update the slice copy, NOT the original slice
	sliceCopy[1] = 7

	fmt.Println("\nupdated sliceCopy:", sliceCopy)
	fmt.Println("original slice:", slice)
	// ↳ the original slice will also be changed because the slice copy is just a pointer
}

// NOTE:
// 	*int     -> indicating the data type of a pointer's value 
// 	*varName -> dereference a pointer to get its value
// 	&varName -> get the memory address of a variable