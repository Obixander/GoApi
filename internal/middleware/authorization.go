package middleware

import (
	"errors"
	"net/http"

	"github.com/Obixander/GoApi/api"
	"github.com/Obixander/GoApi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAutherizedError = errors.New("Invalid username or token.")

func Autherization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Autherization")

		var err error
		if username == "" || token == "" {
			log.Error(UnAutherizedError)
			api.RequestErrorHandler(w, UnAutherizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandling(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAutherizedError)
			api.RequestErrorHandler(w, UnAutherizedError)
			return
		}

		next.ServeHTTP(w, r)

	})
}
