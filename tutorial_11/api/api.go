// Package api defines the request and response interfaces
package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceRequest struct {
	Username string
}

type CoinBalanceResponse struct {
	ResultCode int
	Username string
	Balance    int64
}

type Error struct {
	ResultCode int
	Message    string
}

// a private general error builder function
func buildError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		ResultCode: code,
		Message:    message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	// public wrapper functions
	IllegalRequestHandler = func(w http.ResponseWriter, err error) {
		buildError(w, err.Error(), http.StatusBadRequest)
	}
	SystemErrorHandler = func(w http.ResponseWriter) {
		buildError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)
