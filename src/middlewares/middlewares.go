package middlewares

import (
	"errors"
	"net/http"

	"github.com/RaguSoftEng/inviteservice_api/util"
)

func PrepareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}

}

func SetAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := util.TokenValid(r)
		if err != nil {
			util.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
