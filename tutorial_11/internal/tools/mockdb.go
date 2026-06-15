// Package tools stores all util tools
package tools

import (
	"math/rand"
	"time"
)

// mock data
var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
	},
}

type mockDB struct{}

func (d *mockDB) SetupDatabase() error {
	return nil
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	mockDelay()

	var data = LoginDetails{}
	data, isExists := mockLoginDetails[username]
	if !isExists {
		return nil
	}

	return &data
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	mockDelay()

	var data = CoinDetails{}
	data, isExists := mockCoinDetails[username]
	if !isExists {
		return nil
	}

	return &data
}

func mockDelay() {
	var delay = rand.Float32() * 300
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
