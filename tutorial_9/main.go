package main

import (
	"fmt"
	"time"
)

func main() {
	// channel tutorial
	// channel is like a container that allows storing data from goroutine
	// channel behavior:
	//	1. wait for a goroutine to finish
	// 	2. assign value to the variable

	// SCENE 1: insert 1 value into a channel
	var c chan int = make(chan int)

	fmt.Println("[SCENE 1] starting a goroutine...")

	go insertOneValueIntoChannel(c)
	fmt.Println("<- c:", <-c) // wait for the value to be set inside a channel

	fmt.Println("goroutine completed!")

	// SCENE 2: insert multiple values into a channel, use traditional for loop to access
	var c2 = make(chan int)

	fmt.Println("\n[SCENE 2] starting a goroutine...")

	go insertMultipleValuesIntoChannel(c2)
	for i := 0; i < 5; i++ {
		fmt.Println("<- c2:", <-c2) // wait for the value to be set inside a channel
	}

	fmt.Println("goroutine completed!")

	// SCENE 3: insert multiple values into a channel, use range to access
	var c3 = make(chan int)

	fmt.Println("\n[SCENE 3] starting a goroutine...")

	go insertMultipleValuesIntoChannel2(c3)
	for i := range c3 { // range will assign channel's value inside the var
		fmt.Println("i:", i)
	}

	fmt.Println("goroutine completed!")

	// SCENE 4: insert multiple values into a buffer channel, use range to access
	var c4 = make(chan int, 5) // create a buffer channel with 5 slots

	fmt.Println("\n[SCENE 4] starting a goroutine...")

	go insertMultipleValuesIntoBufferChannel(c4)
	for i := range c4 { // range will assign channel's value inside the var
		fmt.Println("i:", i)
		time.Sleep(1 * time.Second)
		// in this case, the function will finish first
		// 	since the channel has extra buffer to store value.
		// function's log will be printed first before all values are printed
	}

	fmt.Println("goroutine completed!")
}

// this function do not return anything but instead just set a value into the channel
func insertOneValueIntoChannel(c chan int) {
	time.Sleep(time.Duration(1000 * time.Millisecond))
	c <- 123
}

func insertMultipleValuesIntoChannel(c chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(500 * time.Millisecond))
		c <- i
	}
	close(c)
	// if channel is not closed,
	// 	program will keep waiting forever
	// 	and throw deadlock error
}

func insertMultipleValuesIntoChannel2(c chan int) {
	defer close(c) // close channel as soon as method ends
	for i := 0; i < 5; i++ {
		time.Sleep(time.Duration(500 * time.Millisecond))
		c <- i
	}
}

func insertMultipleValuesIntoBufferChannel(c chan int) {
	defer close(c) // close channel as soon as method ends
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("Function insertMultipleValuesIntoBufferChannel() has ended!")
}
