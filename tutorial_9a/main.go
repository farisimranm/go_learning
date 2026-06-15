package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_CHICKEN_PRICE float32 = 5
const MAX_TOFU_PRICE float32 = 3

var websites = []string{"nsk.com", "jayagrocer.my", "faizkimiayam.net"}

func main() {
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendNotification(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(1 * time.Second)
		var price = rand.Float32() * 20
		fmt.Printf("[checkChickenPrices] website: %v, price: %v\n", website, price)
		if price <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(1 * time.Second)
		var price = rand.Float32() * 20
		fmt.Printf("[checkTofuPrices] website: %v, price: %v\n", website, price)
		if price <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

func sendNotification(chickenChannel chan string, tofuChannel chan string) {
	// which channel returns a value first, the case will be executed
	// both channels will race to be the first to have a value
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nSMS SENT - Found the best chicken deal at %v\n", website)
	case website := <-tofuChannel:
		fmt.Printf("\nEMAIL SENT - Found the best tofu deal at %v\n", website)
	}
}
