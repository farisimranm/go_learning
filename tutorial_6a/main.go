package main

import "fmt"

type gasEngine struct {
	kmPerLiter uint
	liters     uint
}

type electricEngine struct {
	kmPerKWH uint
	kwh      uint
}

type engine interface {
	kmLeft() uint
}

// assign this function to gasEngine struct
func (e gasEngine) kmLeft() uint {
	return e.liters * e.kmPerLiter
}

// assign this function to electricEngine struct
func (e electricEngine) kmLeft() uint {
	return e.kwh * e.kmPerKWH
}

// use interface as argument, any struct containing the same method name can be passed in
func canMakeIt(e engine, targetKM uint) bool {
	return targetKM <= e.kmLeft()
}

func main() {
	var targetKM uint = 400

	var kancilEngine = gasEngine{25, 15}

	if canMakeIt(kancilEngine, targetKM) {
		fmt.Println("Kancil can make it with full tank!")
	} else {
		fmt.Println("Kancil cannot make it with full tank, need to refuel first!")
	}

	var bydEngine = electricEngine{40, 30}

	if canMakeIt(bydEngine, targetKM) {
		fmt.Println("BYD can make it with full battery!")
	} else {
		fmt.Println("BYD cannot make it with full battery, need to recharge first!")
	}
}
