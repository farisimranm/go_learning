package main

import (
	"fmt"
	"strings"
)

func main()  {
	// in Go, strings are stored as an array of bytes. every character is a byte with UTF-8 format 
	str := "résumé"
	fmt.Println("str:", str)

	indexedStr := str[0]
	fmt.Printf("value of str[0] is %v with type %T\n", indexedStr, indexedStr)

	for i, v := range str {
		fmt.Printf("i: %v, v: %v\n", i, v)
	}

	// rune
	myRune := 'a'
	fmt.Println("\nmyRune:", myRune)

	// string concatination
	strSlice := []string{"d", "e", "e", "z", "n", "u", "t", "s"}
	catStr := ""
	var sb strings.Builder

	for _, char := range strSlice {
		catStr += char
		sb.WriteString(char)
	}
	fmt.Println("\nstrSlice:", strSlice)
	fmt.Println("catStr:", catStr)
	fmt.Println("sb.toString():", sb.String())
}