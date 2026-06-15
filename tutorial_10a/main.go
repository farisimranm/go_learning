package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type contactInfo struct {
	Name string
	Email string
}

type purchaseInfo struct {
	Name string
	Price float32
	Amount int
}

func main()  {
	// calling a generic method with explicit type because it cant be inferred automatically
	var contactList []contactInfo = loadJSON[contactInfo]("./tutorial_10b/contactInfos.json")
	fmt.Printf("contactList: %+v\n", contactList)
	
	var purchaseList []purchaseInfo = loadJSON[purchaseInfo]("./tutorial_10b/purchaseInfos.json")
	fmt.Printf("purchaseList: %+v\n", purchaseList)
}

func loadJSON[T contactInfo | purchaseInfo](filePath string) []T {
	data, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return make([]T, 0)
	}

	result := []T{}
	json.Unmarshal(data, &result)

	return result
}
