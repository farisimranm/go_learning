package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"

	// our custom packages
	"tutorial_11/api"
	"tutorial_11/internal/tools"
)

var UnauthorizedError = errors.New("Invalid user.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get username from query param
		var username string = r.URL.Query().Get("username")

		// get Authorization from request header
		var token = r.Header.Get("Authorization")

		var err error

		// validate request
		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.IllegalRequestHandler(w, UnauthorizedError)
			return
		}

		// call DB
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.SystemErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil) || (token != (*loginDetails).AuthToken) {
			log.Error(UnauthorizedError)
			api.IllegalRequestHandler(w, UnauthorizedError)
			return
		}

		// proceed to next middleware if any. else return this function
		next.ServeHTTP(w, r)
	})
}
