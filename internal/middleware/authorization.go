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
		var token = r.Header.Get("Authorization")

		var err error
		if username == "" || token == "" {
			log.Error("Username:" + username)
			log.Error("Token:" + token)
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

		var LoginDetails *tools.LoginDetails
		LoginDetails = (*database).GetUserLoginDetails(username)

		if LoginDetails == nil || (token != (*LoginDetails).AuthToken) {
			log.Error(UnAutherizedError)
			api.RequestErrorHandler(w, UnAutherizedError)
			return
		}

		next.ServeHTTP(w, r)

	})
}
