package main

import "fmt"

type car struct {
	color string
	price float64
	year uint
}

type driver struct {
	name string
	carInfo car
	license			// declaring field without field name to merge subfields into parent's fields
}

type license struct {
	hasExpired bool
}

func main()  {
	// 1. unassigned struct
	var myCar car
	fmt.Println("myCar:", myCar)
	fmt.Println("myCar.color:", myCar.color)
	fmt.Println("myCar.price:", myCar.price)
	fmt.Println("myCar.doors:", myCar.year)

	// 2. assigned struct
	var kancil car = car{color: "navyBlue", price: 3500, year: 2004}
	fmt.Println("\nkancil:", kancil)
	fmt.Println("kancil.color:", kancil.color)
	fmt.Println("kancil.price:", kancil.price)
	fmt.Println("kancil.doors:", kancil.year)

	// 3. assigned struct without mentioning attribute names
	var lancer car = car{"red", 315436.10, 2008}

	fmt.Println("\nlancer:", lancer)

	lancer.color = "white" // attribute reassignment

	fmt.Println("lancer.color:", lancer.color)
	fmt.Println("lancer.price:", lancer.price)
	fmt.Println("lancer.doors:", lancer.year)

	// 4. nested struct
	var mihi driver
	mihi.name = "mihi"
	mihi.carInfo = lancer

	fmt.Println("\nmihi:", mihi)
	fmt.Println("mihi.hasExpired:", mihi.hasExpired) 	// now driver.license.hasExpired now becomes driver.hasExpired

	// 5. unreusable struct
	var instantStruct = struct {
		name string
		age uint
	}{"alex", 23}

	fmt.Println("\ninstantStruct:", instantStruct)
	fmt.Println("instantStruct.name:", instantStruct.name)
	fmt.Println("instantStruct.age:", instantStruct.age)
}