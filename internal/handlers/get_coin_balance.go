package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Obixander/GoApi/internal/api"
	"github.com/Obixander/GoApi/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandling(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandling(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: (*&tokenDetails).Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandling(w)
		return
	}
}
