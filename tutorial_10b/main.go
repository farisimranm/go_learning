package main

import (
	"fmt"
)

type windowsOS struct {
	version string
	year int
}

type macOS struct {
	ver string
	manufacturedAt int
}

type desktop[T windowsOS | macOS] struct {
	brand string
	price float32
	os T
}

func main() {
	// using generic "any" in a struct
	var laptop = desktop[windowsOS]{
		brand: "Asus ROG",
		price: 6999.00,
		os: windowsOS{
			version: "Windows 11",
			year: 2026,
		},
	}

	var macbook = desktop[macOS]{
		brand: "MacBook Pro",
		price: 5990.60,
		os: macOS{
			ver: "MacOS 26 Tahoe",
			manufacturedAt: 2025,
		},
	}

	fmt.Printf("laptop: %+v\n", laptop)
	fmt.Printf("macbook: %+v\n", macbook)
}