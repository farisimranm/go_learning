package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"

	// our custom packages
	"tutorial_11/api"
	"tutorial_11/internal/tools"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var coinBalanceRequest = api.CoinBalanceRequest{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&coinBalanceRequest, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.SystemErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		api.SystemErrorHandler(w)
		return
	}

	var coinDetails *tools.CoinDetails = (*database).GetUserCoins(coinBalanceRequest.Username)
	if coinDetails == nil {
		// assuming that the authorization has passed, a nil data is considered as internal system error
		log.Error(err)
		api.SystemErrorHandler(w)
		return
	}

	var coinBalanceResponse = api.CoinBalanceResponse{
		ResultCode: http.StatusOK,
		Username:   coinDetails.Username,
		Balance:    coinDetails.Coins,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(coinBalanceResponse)
	if err != nil {
		log.Error(err)
		api.SystemErrorHandler(w)
		return
	}
}
